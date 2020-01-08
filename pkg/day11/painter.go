package day11

import (
	"fmt"
	"math"
	"strings"

	"github.com/bclement/advent2019/pkg/day9"
)

// Direction indicates where the robot is facing
type Direction int

const (
	north = iota
	south
	east
	west
)

func (dir Direction) String() string {
	switch dir {
	case north:
		return "north"
	case south:
		return "south"
	case east:
		return "east"
	case west:
		return "west"
	default:
		return "invalid"
	}
}

// Color indicates a color for a cell
type Color int64

const (
	black = 0
	white = 1
)

func (c Color) String() string {
	if c == black {
		return "black"
	} else if c == white {
		return "white"
	}
	return "invalid"
}

// Command tells the robot how to move next
type Command int64

const (
	turnLeft  = 0
	turnRight = 1
)

func (c Command) String() string {
	if c == turnLeft {
		return "turn left"
	} else if c == turnRight {
		return "turn right"
	}
	return "invalid"
}
func (dir Direction) applyCommand(c Command) (Direction, error) {
	var leftDir Direction
	var rightDir Direction
	switch dir {
	case north:
		leftDir = west
		rightDir = east
	case south:
		leftDir = east
		rightDir = west
	case east:
		leftDir = north
		rightDir = south
	case west:
		leftDir = south
		rightDir = north
	}
	if c == turnLeft {
		return leftDir, nil
	} else if c == turnRight {
		return rightDir, nil
	}

	return -1, fmt.Errorf("Invalid direction %v", dir)
}

type point struct {
	x, y int
}

type cell struct {
	color   Color
	painted bool
}

// IOHarness uses channels for intcomputer IO
type IOHarness struct {
	in, out chan int64
}

// NewIOHarness creates a new IOHarness object
func NewIOHarness() IOHarness {
	in := make(chan int64, 10)
	out := make(chan int64, 10)
	return IOHarness{in, out}
}

// Close closes the IO harness
func (io IOHarness) Close() {
	close(io.in)
	close(io.out)
}

func (io IOHarness) Read() (i int64, err error) {
	return <-io.in, nil
}

// SendInput sends the integer to the input channel
func (io IOHarness) SendInput(i int64) {
	io.in <- i
}

// Poll attempts to read from input. returns false if no input available
func (io IOHarness) Poll() (int64, bool) {
	select {
	case i := <-io.in:
		return i, true
	default:
		return 0, false
	}
}

func (io IOHarness) Write(i int64) error {
	io.out <- i
	return nil
}

func (io IOHarness) waitForOutput(errc chan error) (Color, Command, bool, error) {
	col, success, err := io.GetOutput(errc)
	if !success {
		return -1, -1, success, err
	}
	com, success, err := io.GetOutput(errc)
	return Color(col), Command(com), success, err
}

// GetOutput gets the next available output from the program
func (io IOHarness) GetOutput(errc chan error) (int64, bool, error) {
	select {
	case value, ok := <-io.out:
		if !ok {
			err := fmt.Errorf("output closed unexpectedly")
			return -1, false, err
		}
		return value, true, nil
	case err, ok := <-errc:
		if !ok {
			err = fmt.Errorf("errc closed unexpectedly")
		}
		return -1, false, err
	}
}

type robot struct {
	comp  *day9.Computer
	cells map[point]*cell
	io    IOHarness
	loc   point
	dir   Direction
}

func newRobot(data []int64) *robot {
	io := NewIOHarness()
	comp := day9.NewComputer(data, io, io, false)
	cells := make(map[point]*cell)
	return &robot{comp, cells, io, point{0, 0}, north}
}

func (r *robot) shutdown() {
	r.io.Close()
}

func (r *robot) paint() (string, error) {
	err := r.run(white)
	if err != nil {
		return "", err
	}
	minx, maxx, miny, maxy := r.getGridDims()
	var sb strings.Builder
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			cell := r.cells[point{x, y}]
			ch := ' '
			if cell != nil && cell.color == white {
				ch = '█'
			}
			sb.WriteRune(ch)
		}
		sb.WriteRune('\n')
	}
	return sb.String(), nil
}

func (r *robot) getGridDims() (minx, maxx, miny, maxy int) {
	minx = math.MaxInt32
	miny = minx
	maxx = math.MinInt32
	maxy = maxx
	for p := range r.cells {
		if minx > p.x {
			minx = p.x
		}
		if maxx < p.x {
			maxx = p.x
		}
		if miny > p.y {
			miny = p.y
		}
		if maxy < p.y {
			maxy = p.y
		}
	}
	return
}
func (r *robot) run(initialColor Color) (err error) {
	errc := make(chan error)
	defer close(errc)
	go r.comp.RunAsync(errc)
	running := true
	r.getCurrentCell().color = initialColor
	for {
		current := r.getCurrentCell()
		//fmt.Printf("currently at %v facing %v, color: %v\n", r.loc, r.dir, current.color)
		r.io.in <- int64(current.color)
		var com Command
		current.color, com, running, err = r.io.waitForOutput(errc)
		if !running {
			return
		}
		//fmt.Printf("got command to paint cell %v, then %v\n", current.color, com)
		current.painted = true
		r.dir, err = r.dir.applyCommand(com)
		if err != nil {
			return
		}
		err = r.move()
		if err != nil {
			return
		}
	}
}

func (r *robot) getCurrentCell() *cell {
	c := r.cells[r.loc]
	if c == nil {
		c = &cell{black, false}
		r.cells[r.loc] = c
	}
	return c
}

func (r *robot) move() error {
	switch r.dir {
	case north:
		r.loc = point{r.loc.x, r.loc.y - 1}
	case south:
		r.loc = point{r.loc.x, r.loc.y + 1}
	case east:
		r.loc = point{r.loc.x + 1, r.loc.y}
	case west:
		r.loc = point{r.loc.x - 1, r.loc.y}
	default:
		return fmt.Errorf("Invalid direction: %v", r.dir)
	}
	return nil
}

func (r *robot) countPainted() int {
	rval := 0
	for _, c := range r.cells {
		if c.painted {
			rval++
		}
	}
	return rval
}

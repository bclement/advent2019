package day13

import (
	"github.com/bclement/advent2019/pkg/day9"
)

var scorePoint = point{-1, 0}

type point struct {
	x, y int64
}

type tile int64

const (
	empty  = 0
	wall   = 1
	block  = 2
	paddle = 3
	ball   = 4
)

type joystickPos int64

const (
	left    = -1
	neutral = 0
	right   = 1
)

type outputType int

const (
	xCoord   = 0
	yCoord   = 1
	tileType = 2
)

type ioHarness struct {
	tiles      map[point]tile
	posInput   joystickPos
	nextOutput outputType
	score      int64
	ballPos    point
	paddlePos  point
	current    point
}

func newIoHarness() *ioHarness {
	return &ioHarness{tiles: make(map[point]tile)}
}

func (io *ioHarness) Read() (int64, error) {
	return int64(io.posInput), nil
}

func (io *ioHarness) Write(i int64) error {
	switch io.nextOutput {
	case xCoord:
		io.current.x = i
		io.nextOutput = yCoord
	case yCoord:
		io.current.y = i
		io.nextOutput = tileType
	case tileType:
		if io.current == scorePoint {
			io.score = i
		} else {
			io.tiles[io.current] = tile(i)
			if i == ball {
				io.ballPos = io.current
			} else if i == paddle {
				io.paddlePos = io.current
			}
			if io.ballPos.x == io.paddlePos.x {
				io.posInput = neutral
			} else if io.ballPos.x < io.paddlePos.x {
				io.posInput = left
			} else {
				io.posInput = right
			}
		}
		io.nextOutput = xCoord
	}
	return nil
}

type arcade struct {
	comp *day9.Computer
	io   *ioHarness
}

func newArcade(fname string, coins int64) (*arcade, error) {
	data, err := day9.ReadDataFile(fname)
	if err != nil {
		return nil, err
	}
	if coins != 0 {
		data[0] = coins
	}
	io := newIoHarness()
	comp := day9.NewComputer(data, io, io, false)
	return &arcade{comp, io}, nil
}

func (a *arcade) run() (err error) {
	return a.comp.Run()
}

func (a *arcade) countBlocks() int {
	rval := 0
	for _, val := range a.io.tiles {
		if val == block {
			rval++
		}
	}
	return rval
}

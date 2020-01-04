package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type position struct {
	vals [3]int
}

func newPos(x, y, z int) position {
	return position{[3]int{x, y, z}}
}

func (p position) String() string {
	return fmt.Sprintf("<x=%v, y=%v, z=%v>", p.vals[0], p.vals[1], p.vals[2])
}

func (p position) equals(other position) bool {
	for i, val := range p.vals {
		if val != other.vals[i] {
			return false
		}
	}
	return true
}

func (p position) axisEquals(other position, axis int) bool {
	return p.vals[axis] == other.vals[axis]
}

type velocity struct {
	vals [3]int
}

func newVel(x, y, z int) velocity {
	return velocity{[3]int{x, y, z}}
}

func (v velocity) String() string {
	return fmt.Sprintf("<x=%v, y=%v, z=%v>", v.vals[0], v.vals[1], v.vals[2])
}

func (v velocity) equals(other velocity) bool {
	for i, val := range v.vals {
		if val != other.vals[i] {
			return false
		}
	}
	return true
}

func (v velocity) axisEquals(other velocity, axis int) bool {
	return v.vals[axis] == other.vals[axis]
}

type moon struct {
	pos *position
	vel *velocity
}

func (m *moon) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", *m.pos, *m.vel)
}

func (m *moon) equals(other *moon) bool {
	return m.pos.equals(*other.pos) && m.vel.equals(*other.vel)
}

func (m *moon) axisEquals(other *moon, axis int) bool {
	return m.pos.axisEquals(*other.pos, axis) && m.vel.axisEquals(*other.vel, axis)
}

func newMoon(pos position) *moon {
	vel := newVel(0, 0, 0)
	return &moon{&pos, &vel}
}

func (m *moon) gravitate(other *moon) {
	for i, val := range m.pos.vals {
		m.vel.vals[i] += relate(val, other.pos.vals[i])
	}
}

func relate(this, other int) int {
	if this > other {
		return -1
	}
	if this < other {
		return 1
	}
	return 0
}

func (m *moon) move() {
	for i, val := range m.vel.vals {
		m.pos.vals[i] += val
	}
}

func (m *moon) getEnergy() int {
	pot := addAbs(m.pos.vals)
	kin := addAbs(m.vel.vals)
	return pot * kin
}

func addAbs(vals [3]int) int {
	fx, fy, fz := float64(vals[0]), float64(vals[1]), float64(vals[2])
	return int(math.Abs(fx) + math.Abs(fy) + math.Abs(fz))
}

func run(fname string, steps int) (int, error) {
	moons, err := read(fname)
	if err != nil {
		return -1, err
	}
	for i := 0; i < steps; i++ {
		fmt.Printf("After %v steps\n", i)
		for _, m := range moons {
			fmt.Printf("%v\n", m)
		}
		for a, m := range moons {
			for b, o := range moons {
				if a == b {
					continue
				}
				m.gravitate(o)
			}
		}
		for _, m := range moons {
			m.move()
		}
		fmt.Println("")
	}
	fmt.Printf("After %v steps\n", steps)
	for _, m := range moons {
		fmt.Printf("%v\n", m)
	}
	rval := 0
	for _, m := range moons {
		rval += m.getEnergy()
	}
	return rval, nil
}

func copyMoons(moons []*moon) []*moon {
	rval := make([]*moon, len(moons))
	for i, m := range moons {
		rval[i] = newMoon(*m.pos)
	}
	return rval
}

func equal(moons, originals []*moon) bool {
	for i, m := range moons {
		if !m.equals(originals[i]) {
			return false
		}
	}
	return true
}

func axisEqual(moons, originals []*moon, axis int) bool {
	for i, m := range moons {
		if !m.axisEquals(originals[i], axis) {
			return false
		}
	}
	return true
}

func getAxisCycle(moons, originals []*moon, axis int) int64 {
	rval := int64(0)
	for !axisEqual(moons, originals, axis) || rval == 0 {
		for a, m := range moons {
			for b, o := range moons {
				if a == b {
					continue
				}
				m.vel.vals[axis] += relate(m.pos.vals[axis], o.pos.vals[axis])
			}
		}
		for _, m := range moons {
			m.pos.vals[axis] += m.vel.vals[axis]
		}
		rval++
	}
	return rval
}

func gcd(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int64, others ...int64) int64 {
	rval := x * y / gcd(x, y)
	for _, val := range others {
		rval = lcm(rval, val)
	}
	return rval
}

func getCycle(fname string) (int64, error) {
	moons, err := read(fname)
	if err != nil {
		return -1, err
	}
	cycles := [3]int64{0, 0, 0}
	for i := 0; i < len(cycles); i++ {
		clones := copyMoons(moons)
		cycles[i] = getAxisCycle(clones, moons, i)
	}
	return lcm(cycles[0], cycles[1], cycles[2]), nil
}

func read(fname string) ([]*moon, error) {
	in, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	re := regexp.MustCompile("^<x=(-?[0-9]+), y=(-?[0-9]+), z=(-?[0-9]+)>.*$")
	var rval []*moon
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		groups := re.FindStringSubmatch(line)
		if groups == nil {
			continue
		}
		fmt.Printf("groups: %v\n", groups)
		x, err := strconv.Atoi(groups[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(groups[2])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(groups[3])
		if err != nil {
			return nil, err
		}
		rval = append(rval, newMoon(newPos(x, y, z)))
	}
	return rval, nil
}

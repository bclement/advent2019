package day10

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func (p point) dist(other point) float64 {
	return findDist(float64(p.x-other.x), float64(p.y-other.y))
}

func (p point) findAngle(other point) (angle float64) {
	dx := p.x - other.x
	dy := p.y - other.y
	if dx == 0 {
		if dy < 0 {
			angle = 180
		} else {
			angle = 0
		}
	} else if dy == 0 {
		if dx < 0 {
			angle = 90
		} else {
			angle = 270
		}
	} else {
		angle = findAngle(float64(dx), float64(dy))
	}
	return
}

func findAngle(dx, dy float64) (angle float64) {
	dx, dy, baseAngle := normalize(dx, dy)
	return baseAngle + ((180 / math.Pi) * math.Atan(dy/dx))
}

func findDist(dx, dy float64) (dist float64) {
	dx = math.Abs(dx)
	dy = math.Abs(dy)
	if dx == 0 {
		dist = dy
	} else if dy == 0 {
		dist = dx
	} else {
		dist = math.Hypot(dx, dy)
	}
	return
}

// make angle calculations the same for all four quadrants
// x increases left to right
// y increases top to bottom (opposite of standard graph)
// after normalization dx is the base of the triangle and dy is the height
// the height of the triangle is always along an axis
// the baseAngle should be added to the angle calculated between the height and hypotenuse
// here is the layout of the four quadrants:
//           DA
//           CB
// 0 degrees is the axis between D and A and angle increases clockwise with 90 degrees
// between AB, 180 between BC and 270 between CD
func normalize(dx, dy float64) (normalDx, normalDy, baseAngle float64) {
	if dx > 0 {
		if dy > 0 {
			// quadrant D
			normalDx, normalDy, baseAngle = dx, dy, 270
		} else {
			// quadtrant C
			normalDx, normalDy, baseAngle = math.Abs(dy), dx, 180
		}
	} else {
		if dy > 0 {
			// quadrant A
			normalDx, normalDy, baseAngle = dy, math.Abs(dx), 0
		} else {
			// quadrant B
			normalDx, normalDy, baseAngle = math.Abs(dx), math.Abs(dy), 90
		}
	}
	return
}

type relation struct {
	p     point
	dist  float64
	angle float64
}

type byAngleDist []relation

func (s byAngleDist) Len() int {
	return len(s)
}

func (s byAngleDist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byAngleDist) Less(i, j int) bool {
	ai := s[i].angle
	aj := s[j].angle
	if ai < aj {
		return true
	} else if ai > aj {
		return false
	} else {
		di := s[i].dist
		dj := s[j].dist
		return di < dj
	}
}

func getRemovalOrder(origin point, points []point) []point {
	rels := getRelations(origin, points, true)
	var rval []point
	current := rels
	var kept []relation
	for len(rval) < len(rels) {
		angle := float64(-1)
		for _, rel := range current {
			if rel.angle != angle {
				angle = rel.angle
				rval = append(rval, rel.p)
				fmt.Printf("%v: Removing %v at angle %v dist %v\n", len(rval), rel.p, rel.angle, rel.dist)
			} else {
				kept = append(kept, rel)
			}
		}
		current = kept
		kept = nil
	}
	return rval
}

func getOptimalPoint(points []point) (best point, count int) {
	count = math.MinInt32
	for _, p := range points {
		rels := getRelations(p, points, false)
		num := len(rels)
		if num > count {
			best = p
			count = num
		}
	}
	return
}

func getRelations(origin point, points []point, includeBlocked bool) []relation {
	var rels []relation
	for _, other := range points {
		dist := origin.dist(other)
		if dist > 0 {
			angle := origin.findAngle(other)
			rels = append(rels, relation{other, dist, angle})
		}
	}
	sort.Sort(byAngleDist(rels))
	if includeBlocked {
		return rels
	}
	return removeBlocked(rels)
}

// takes in relations sorted first by angle then by distance ascending
// returns list with only closest relation for each unique angle
func removeBlocked(rels []relation) []relation {
	var rval []relation
	angle := float64(-1)
	for _, r := range rels {
		if r.angle != angle {
			rval = append(rval, r)
		}
		angle = r.angle
	}
	return rval
}

func readInput(fname string) ([]point, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var rval []point
	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			if c == '#' {
				rval = append(rval, point{x, y})
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return rval, err
	}
	return rval, nil
}

package day3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	up = iota
	down
	left
	right
)

type step struct {
	dir   direction
	count int
}

type point struct {
	x, y int
}

func findShortestPath(points map[point]bool, wire1Steps, wire2steps []step) (point, int) {
	var shortest point
	minDist := math.MaxInt32
	for p := range points {
		dist := calculatePathLen(p, wire1Steps)
		dist += calculatePathLen(p, wire2steps)
		if dist < minDist {
			shortest = p
			minDist = dist
		}
	}
	return shortest, minDist
}

func calculatePathLen(target point, steps []step) int {
	count := 0
	forEachPoint(steps, func(p point) bool {
		count++
		return p != target
	})
	return count
}

func findNearestPoint(points map[point]bool, target point) (point, int) {
	var nearest point
	minDist := math.MaxInt32
	for p := range points {
		dist := calculateDistance(target, p)
		if dist < minDist {
			minDist = dist
			nearest = p
		}
	}
	return nearest, minDist
}

func calculateDistance(a, b point) int {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))
	return int(dx + dy)
}

func getIntersections(wire1Steps, wire2Steps []step) map[point]bool {
	wire1Points := collectPoints(wire1Steps)
	return getCommonPoints(wire2Steps, wire1Points)
}

func getCommonPoints(steps []step, prev map[point]bool) map[point]bool {
	set := make(map[point]bool)
	forEachPoint(steps, func(p point) bool {
		if prev[p] {
			set[p] = true
		}
		return true
	})
	return set
}

func collectPoints(steps []step) map[point]bool {
	set := make(map[point]bool)
	forEachPoint(steps, func(p point) bool {
		set[p] = true
		return true
	})
	return set
}

func forEachPoint(steps []step, c func(point) bool) {
	x := 0
	y := 0
	for _, step := range steps {
		var axis *int
		var neg bool
		switch step.dir {
		case up:
			axis = &y
			neg = false
		case down:
			axis = &y
			neg = true
		case left:
			axis = &x
			neg = true
		case right:
			axis = &x
			neg = false
		}
		for i := 0; i < step.count; i++ {
			if neg {
				*axis--
			} else {
				*axis++
			}
			if !c(point{x, y}) {
				return
			}
		}
	}
}

func readInput() ([]step, []step, error) {
	f, err := os.Open("../../resources/day3/input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line1, err := readLine(scanner)
	if err != nil {
		return nil, nil, err
	}
	line2, err := readLine(scanner)
	if err != nil {
		return nil, nil, err
	}
	return line1, line2, nil
}

func readLine(scanner *bufio.Scanner) ([]step, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("Expected line of input")
	}
	line := scanner.Text()
	parts := strings.Split(line, ",")
	rval := make([]step, len(parts))
	for i := 0; i < len(parts); i++ {
		step, err := parseStep(parts[i])
		if err != nil {
			return nil, err
		}
		rval[i] = *step
	}
	return rval, nil
}

func parseStep(text string) (*step, error) {
	flag := text[:1]
	dir, err := parseDir(flag)
	if err != nil {
		return nil, err
	}
	count, err := strconv.Atoi(text[1:])
	if err != nil {
		return nil, err
	}
	return &step{dir, count}, nil
}

func parseDir(flag string) (direction, error) {
	var rval direction
	switch flag {
	case "U":
		rval = up
	case "D":
		rval = down
	case "R":
		rval = right
	case "L":
		rval = left
	default:
		return up, fmt.Errorf("Unknown direction flag: %v", flag)
	}
	return rval, nil
}

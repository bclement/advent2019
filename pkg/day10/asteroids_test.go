package day10

import (
	"math"
	"testing"
)

func TestAngles(t *testing.T) {
	center := point{0, 0}
	n := point{0, -1}
	assertAngle(t, center, n, 0)
	s := point{0, 1}
	assertAngle(t, center, s, 180)
	e := point{1, 0}
	assertAngle(t, center, e, 90)
	w := point{-1, 0}
	assertAngle(t, center, w, 270)
	nw := point{-1, -1}
	assertAngle(t, center, nw, 315)
	ne := point{1, -1}
	assertAngle(t, center, ne, 45)
	se := point{1, 1}
	assertAngle(t, center, se, 135)
	sw := point{-1, 1}
	assertAngle(t, center, sw, 225)
}

func assertAngle(t *testing.T, center, other point, expected float64) {
	result := center.findAngle(other)
	if math.Abs(result-expected) > 0.00001 {
		t.Errorf("Angle between %v and %v: expected %v, got %v", center, other, expected, result)
	}
}

func TestSmall(t *testing.T) {
	assertBest(t, "../../resources/day10/small.txt", point{3, 4}, 8)
}

func TestMed(t *testing.T) {
	assertBest(t, "../../resources/day10/1-2-35.txt", point{1, 2}, 35)
}

func TestFull(t *testing.T) {
	assertBest(t, "../../resources/day10/input.txt", point{19, 11}, 230)
}

func TestFullRemoval(t *testing.T) {
	assertRemoval(t, "../../resources/day10/input.txt", point{19, 11}, point{12, 5}, 200)
}

func TestSmallRemoval(t *testing.T) {
	assertRemoval(t, "../../resources/day10/rem-small.txt", point{8, 3}, point{14, 3}, 36)
}

func TestLargeRemoval(t *testing.T) {
	assertRemoval(t, "../../resources/day10/11-13-210.txt", point{11, 13}, point{8, 2}, 200)
}

func assertRemoval(t *testing.T, fname string, origin, expected point, nth int) {
	points, err := readInput(fname)
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	result := getRemovalOrder(origin, points)
	target := result[nth-1]
	if target != expected {
		t.Fatalf("Expected point %v got %v", expected, target)
	}
}

func assertBest(t *testing.T, fname string, expected point, expectedCount int) {
	points, err := readInput(fname)
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	best, count := getOptimalPoint(points)
	if best != expected {
		t.Errorf("Expected point %v got %v", expected, best)
	}
	if count != expectedCount {
		t.Fatalf("Expected count %v got %v", expectedCount, count)
	}
}

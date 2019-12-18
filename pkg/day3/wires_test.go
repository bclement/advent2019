package day3

import "testing"

func TestInput(t *testing.T) {
	wire1, wire2, err := readInput()
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	assertStep(t, wire1[0], step{right, 999})
	assertStep(t, wire2[0], step{left, 996})
}

func assertStep(t *testing.T, actual, expected step) {
	if actual.dir != expected.dir {
		t.Errorf("Expected dir %v, got %v", expected.dir, actual.dir)
	}
	if actual.count != expected.count {
		t.Errorf("Expected count %v, got %v", expected.count, actual.count)
	}
}

func TestCollectPoints(t *testing.T) {
	steps := make([]step, 4)
	steps[0] = step{right, 1}
	steps[1] = step{up, 2}
	steps[2] = step{left, 3}
	steps[3] = step{down, 4}
	results := collectPoints(steps)
	expected := []point{point{1, 0}, point{1, 1}, point{1, 2},
		point{0, 2}, point{-1, 2}, point{-2, 2}, point{-2, 1},
		point{-2, 0}, point{-2, -1}, point{-2, -2}}
	if len(results) != len(expected) {
		t.Fatalf("Expected length %v, got %v", len(expected), len(results))
	}
	for _, point := range expected {
		if !results[point] {
			t.Errorf("Missing point %v", point)
		}
	}
}

func TestGetIntersections(t *testing.T) {
	wire1Steps := make([]step, 2)
	wire1Steps[0] = step{right, 5}
	wire1Steps[1] = step{up, 10}
	wire2Steps := make([]step, 2)
	wire2Steps[0] = step{up, 5}
	wire2Steps[1] = step{right, 10}
	results := getIntersections(wire1Steps, wire2Steps)
	if len(results) != 1 {
		t.Fatalf("Expected length of 1, got %v", len(results))
	}
	if !results[point{5, 5}] {
		t.Errorf("Unexpected results: %v", results)
	}
}

func TestFindNearestPoint(t *testing.T) {
	points := make(map[point]bool)
	points[point{100, 0}] = true
	points[point{10, 10}] = true
	result, _ := findNearestPoint(points, point{0, 0})
	if result.x != 10 || result.y != 10 {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestFindNearestIntersection(t *testing.T) {
	wire1, wire2, err := readInput()
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	points := getIntersections(wire1, wire2)
	_, minDist := findNearestPoint(points, point{0, 0})
	expected := 8015
	if minDist != expected {
		t.Errorf("Expected min dist of %v, got %v", expected, minDist)
	}
}

func TestFindShortestPath(t *testing.T) {
	wire1, wire2, err := readInput()
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	points := getIntersections(wire1, wire2)
	_, minDist := findShortestPath(points, wire1, wire2)
	expected := 163676
	if minDist != expected {
		t.Errorf("Expected min dist of %v, got %v", expected, minDist)
	}
}

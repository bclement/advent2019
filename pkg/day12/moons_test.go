package day12

import "testing"

func TestSmall(t *testing.T) {
	assert(t, "../../resources/day12/small.txt", 10, 179)
}

func TestMed(t *testing.T) {
	assert(t, "../../resources/day12/medium.txt", 100, 1940)
}

func TestFull(t *testing.T) {
	assert(t, "../../resources/day12/input.txt", 1000, 12644)
}

func assert(t *testing.T, fname string, steps, expected int) {
	result, err := run(fname, steps)
	if err != nil {
		t.Fatalf("Problem running simulation: %v", err)
	}
	if result != expected {
		t.Fatalf("Expected total of %v, got %v", expected, result)
	}
}

func TestSmallRepeat(t *testing.T) {
	assertRepeat(t, "../../resources/day12/small.txt", 2772)
}

func TestMediumRepeat(t *testing.T) {
	assertRepeat(t, "../../resources/day12/medium.txt", 4686774924)
}

func TestFullRepeat(t *testing.T) {
	assertRepeat(t, "../../resources/day12/input.txt", 290314621566528)
}

func assertRepeat(t *testing.T, fname string, expected int64) {
	result, err := getCycle(fname)
	if err != nil {
		t.Fatalf("Problem running simulation: %v", err)
	}
	if result != expected {
		t.Fatalf("Expected %v steps, got %v", expected, result)
	}
}

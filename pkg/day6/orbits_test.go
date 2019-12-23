package day6

import (
	"os"
	"testing"
)

func TestSmallSet(t *testing.T) {
	assertRun(t, "../../resources/day6/small.txt", 42)
}

func TestFullSet(t *testing.T) {
	assertRun(t, "../../resources/day6/input.txt", 295834)
}

func assertRun(t *testing.T, file string, expected int) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("Problem opening input: %v", err)
	}
	defer f.Close()
	result, err := run(f)
	if err != nil {
		t.Fatalf("Problem during run: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSmallJumpSet(t *testing.T) {
	assertJumps(t, "../../resources/day6/jumps.txt", "YOU", "SAN", 4)
}

func TestFullJumpSet(t *testing.T) {
	assertJumps(t, "../../resources/day6/input.txt", "YOU", "SAN", 361)
}

func assertJumps(t *testing.T, file, from, to string, expected int) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatalf("Problem opening input: %v", err)
	}
	defer f.Close()
	result, err := getNumJumps(f, from, to)
	if err != nil {
		t.Fatalf("Problem during run: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

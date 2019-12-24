package day7

import (
	"fmt"
	"testing"

	"github.com/bclement/advent2019/pkg/day5"
)

func TestSmall(t *testing.T) {
	data := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	settings := []int{4, 3, 2, 1, 0}
	assertRun(t, data, settings, 43210)
	data = []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	settings = []int{0, 1, 2, 3, 4}
	assertRun(t, data, settings, 54321)
	data = []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
		1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	settings = []int{1, 0, 4, 3, 2}
	assertRun(t, data, settings, 65210)
}

func TestSmallLoop(t *testing.T) {
	data := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
		27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	assertFindMax(t, data, 139629729, 5, 9)
}

func TestFindMax(t *testing.T) {
	data := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	assertFindMax(t, data, 43210, 0, 4)
}

func TestFull(t *testing.T) {
	data, err := day5.ReadDataFile("../../resources/day7/input.txt")
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	assertFindMax(t, data, 13848, 0, 4)
}

func TestFullLoop(t *testing.T) {
	data, err := day5.ReadDataFile("../../resources/day7/input.txt")
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	assertFindMax(t, data, 12932154, 5, 9)
}

func assertFindMax(t *testing.T, data []int, expected, min, max int) {
	settings, result, err := findMax(data, min, max)
	if err != nil {
		t.Fatalf("Problem running test: %v", err)
	}
	fmt.Printf("Settings: %v\n", settings)
	if result != expected {
		t.Fatalf("Expected max %v, got %v", expected, result)
	}
}

func assertRun(t *testing.T, data, settings []int, expected int) {
	result, err := run(data, settings)
	if err != nil {
		t.Fatalf("Problem running test: %v", err)
	}
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

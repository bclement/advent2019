package day11

import (
	"testing"

	"github.com/bclement/advent2019/pkg/day9"
)

func TestCount(t *testing.T) {
	data, err := day9.ReadDataFile("../../resources/day11/input.txt")
	if err != nil {
		t.Fatalf("Problem reading data: %v", err)
	}
	r := newRobot(data)
	defer r.shutdown()
	err = r.run(black)
	if err != nil {
		t.Fatalf("Problem running robot: %v", err)
	}
	expected := 2293
	actual := r.countPainted()
	if actual != expected {
		t.Fatalf("Expected count of %v, got %v", expected, actual)
	}
}

func TestPaint(t *testing.T) {
	data, err := day9.ReadDataFile("../../resources/day11/input.txt")
	if err != nil {
		t.Fatalf("Problem reading data: %v", err)
	}
	r := newRobot(data)
	defer r.shutdown()
	result, err := r.paint()
	if err != nil {
		t.Fatalf("Problem running robot: %v", err)
	}
	t.Fatalf("\n%v", result)
}

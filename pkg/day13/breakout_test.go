package day13

import "testing"

func TestTileCount(t *testing.T) {
	a, err := newArcade("../../resources/day13/input.txt", 0)
	if err != nil {
		t.Fatalf("Problem creating arcade: %v", err)
	}
	err = a.run()
	if err != nil {
		t.Fatalf("Problem running arcade: %v", err)
	}
	result := a.countBlocks()
	expected := 260
	if result != expected {
		t.Errorf("Expected %v blocks got %v", expected, result)
	}
}

func TestPlay(t *testing.T) {
	a, err := newArcade("../../resources/day13/input.txt", 2)
	if err != nil {
		t.Fatalf("Problem creating arcade: %v", err)
	}
	err = a.run()
	if err != nil {
		t.Fatalf("Problem running arcade: %v", err)
	}
	expected := int64(12952)
	if a.io.score != expected {
		t.Errorf("Expected %v score got %v", expected, a.io.score)
	}
}

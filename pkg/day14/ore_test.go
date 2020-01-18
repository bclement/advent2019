package day14

import "testing"

func TestA(t *testing.T) {
	assertOreReq(t, "../../resources/day14/a.txt", 31)
}

func TestB(t *testing.T) {
	assertOreReq(t, "../../resources/day14/b.txt", 165)
}

func TestC(t *testing.T) {
	assertOreReq(t, "../../resources/day14/c.txt", 13312)
}

func TestInput(t *testing.T) {
	assertOreReq(t, "../../resources/day14/input.txt", 469536)
}

func assertOreReq(t *testing.T, fname string, expected int64) {
	fact, err := newFactory(fname)
	if err != nil {
		t.Fatalf("Problem building factory: %v", err)
	}
	result, err := fact.getOreRequirement(1)
	if err != nil {
		t.Fatalf("Problem running ore calc: %v", err)
	}
	if result != expected {
		t.Fatalf("Expected ore count of %v got %v", expected, result)
	}
}

func TestMaxC(t *testing.T) {
	assertMaxFuel(t, "../../resources/day14/c.txt", 82892753)
}

func TestMaxE(t *testing.T) {
	assertMaxFuel(t, "../../resources/day14/e.txt", 460664)
}

func TestMaxInput(t *testing.T) {
	assertMaxFuel(t, "../../resources/day14/input.txt", 3343477)
}

func assertMaxFuel(t *testing.T, fname string, expected int64) {
	fact, err := newFactory(fname)
	if err != nil {
		t.Fatalf("Problem building factory: %v", err)
	}
	maxFuel, err := fact.getMaxFuel(1000000000000)
	if err != nil {
		t.Fatalf("Error calculating max fuel: %v", err)
	}
	if maxFuel != expected {
		t.Fatalf("Expected max fuel of %v got %v", expected, maxFuel)
	}
}

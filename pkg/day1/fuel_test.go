package day1

import (
	"fmt"
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	in, err := getInput()
	if err != nil {
		t.Errorf("Problem opening input: %v", err)
		return
	}
	defer in.Close()
	reader := NewIntLineReader(in)
	count := 0
	for reader.HasNext() {
		i, err := reader.Next()
		if err != nil {
			t.Errorf("Problem parsing int: %v", err)
			return
		}
		fmt.Printf("%v\n", i)
		count++
	}
	if count != 100 {
		t.Errorf("Expected 100 numbers, got %v ", count)
	}
}

func getInput() (*os.File, error) {
	return os.Open("../../resources/day1/input.txt")
}

func TestFuelCalc(t *testing.T) {
	assertFuelCalc(t, 12, 2)
	assertFuelCalc(t, 14, 2)
	assertFuelCalc(t, 1969, 654)
	assertFuelCalc(t, 100756, 33583)
}

func assertFuelCalc(t *testing.T, mass int, expected int) {
	assertCalc(t, calculateFuel, mass, expected)
}

func assertCalc(t *testing.T, f func(int) int, mass int, expected int) {
	res := f(mass)
	if res != expected {
		t.Errorf("Expected %v from mass cal of %v, got %v", expected, mass, res)
	}
}

func TestTotalFuel(t *testing.T) {
	assertTotalFuel(t, false, 3246455)
}

func assertTotalFuel(t *testing.T, addFuelFuel bool, expected int) {
	in, err := getInput()
	if err != nil {
		t.Errorf("Problem opening input: %v", err)
		return
	}
	defer in.Close()
	total, err := calculateTotalFuel(in, addFuelFuel)
	if err != nil {
		t.Errorf("Problem calculating fuel: %v", err)
		return
	}
	if total != expected {
		t.Errorf("Expected total of %v got %v", expected, total)
		return
	}
}

func TestModuelFuel(t *testing.T) {
	assertModuleCalc(t, 14, 2)
	assertModuleCalc(t, 1969, 966)
	assertModuleCalc(t, 100756, 50346)
}

func assertModuleCalc(t *testing.T, mass int, expected int) {
	assertCalc(t, calculateModuleFuel, mass, expected)
}

func TestTotalModuleFuel(t *testing.T) {
	assertTotalFuel(t, true, 4866824)
}

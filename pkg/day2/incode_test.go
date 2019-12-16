package day2

import (
	"testing"
)

func TestInput(t *testing.T) {
	result, err := readInput("../../resources/day2/input.txt")
	if err != nil {
		t.Errorf("Problem reading input: %v", err)
		return
	}
	if result[0] != 1 {
		t.Errorf("Expected first number to be 1, but was %v", result[0])
		return
	}
	last := result[len(result)-1]
	if last != 0 {
		t.Errorf("Expected last number to be 0 but was %v", last)
	}
}

func TestRun(t *testing.T) {
	data := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}
	assertRun(t, data, expected)
	data = []int{2, 3, 0, 3, 99}
	expected = []int{2, 3, 0, 6, 99}
	assertRun(t, data, expected)
	data = []int{2, 4, 4, 5, 99, 0}
	expected = []int{2, 4, 4, 5, 99, 9801}
	assertRun(t, data, expected)
	data = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	assertRun(t, data, expected)
}

func assertRun(t *testing.T, data, expected []int) {
	_, err := run(data)
	if err != nil {
		t.Errorf("Problem during run: %v", err)
		return
	}
	for i := 0; i < len(data); i++ {
		if data[i] != expected[i] {
			t.Errorf("Expected value %v at index %v got %v", expected[i], i, data[i])
		}
	}
}

func TestIntcode(t *testing.T) {
	result, err := runWithNounVerb(12, 2)
	if err != nil {
		t.Errorf("Problem running input: %v", err)
		return
	}
	if result != 5290681 {
		t.Errorf("Unexpected output: %v", result)
	}
}

func TestBruteForce(t *testing.T) {
	assertNounVerb(t, 5290681, 12, 2)
	assertNounVerb(t, 19690720, 57, 41)
}

func assertNounVerb(t *testing.T, answer, expectedNoun, expectedVerb int) {
	noun, verb, err := findNounVerb(answer)
	if err != nil {
		t.Errorf("Problem: %v", err)
		return
	}
	if noun != expectedNoun || verb != expectedVerb {
		t.Errorf("Got noun/verb %v/%v expected %v/%v", noun, verb, expectedNoun, expectedVerb)
	}
}

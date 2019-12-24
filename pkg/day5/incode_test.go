package day5

import (
	"fmt"
	"reflect"
	"testing"
)

type slicesrc struct {
	input []int
	pos   int
}

func (src *slicesrc) Read() (int, error) {
	if src.pos > len(src.input)-1 {
		return -1, fmt.Errorf("Ran out of input at %v", src.pos)
	}
	rval := src.input[src.pos]
	src.pos++
	return rval, nil
}

type slicedest struct {
	output []int
}

func (dest *slicedest) Write(i int) error {
	dest.output = append(dest.output, i)
	return nil
}

func TestParseInst(t *testing.T) {
	data := []int{1002, 4, 3, 4, 33}
	result, err := parseInstruction(0, data)
	if err != nil {
		t.Fatalf("Problem parsing input: %v", err)
	}
	if result.opcode != mult {
		t.Errorf("Expected mult, got %v", result.opcode)
	}
	expectedModes := []int{0, 1, 0}
	expectedArgs := []int{4, 3, 4}
	if !reflect.DeepEqual(expectedModes, result.modes) {
		t.Errorf("Expected modes %v, got %v", expectedModes, result.modes)
	}
	if !reflect.DeepEqual(expectedArgs, result.arguments) {
		t.Errorf("Expected args %v, got %v", expectedArgs, result.arguments)
	}
}

func TestInOut(t *testing.T) {
	in := []int{42}
	data := []int{3, 0, 4, 0, 99}
	assertRunData(t, data, in, in)
}

func TestPosEq(t *testing.T) {
	in := []int{8}
	data := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := []int{1}
	assertRunData(t, data, in, expected)
}

func TestPosNEq(t *testing.T) {
	data := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int{42}
	expected := []int{0}
	assertRunData(t, data, in, expected)
}

func TestPosLT(t *testing.T) {
	data := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int{7}
	expected := []int{1}
	assertRunData(t, data, in, expected)
}

func TestPosGTE(t *testing.T) {
	data := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int{9}
	expected := []int{0}
	assertRunData(t, data, in, expected)
}

func TestImmEq(t *testing.T) {
	data := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	in := []int{8}
	expected := []int{1}
	assertRunData(t, data, in, expected)
}

func TestPJump(t *testing.T) {
	data := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	in := []int{0}
	expected := []int{0}
	assertRunData(t, data, in, expected)
}

func TestPJumpNE(t *testing.T) {
	data := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	in := []int{42}
	expected := []int{1}
	assertRunData(t, data, in, expected)
}

func TestRun1(t *testing.T) {
	input := []int{1}
	expected := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 13285749}
	assertRunFullData(t, input, expected)
}

func TestRun5(t *testing.T) {
	input := []int{5}
	expected := []int{5000972}
	assertRunFullData(t, input, expected)
}

func assertRunFullData(t *testing.T, input, expected []int) {
	data, err := readData()
	if err != nil {
		t.Fatalf("Problem reading data: %v", err)
	}
	assertRunData(t, data, input, expected)
}

func assertRunData(t *testing.T, data, input, expected []int) {
	in := &slicesrc{input, 0}
	out := &slicedest{nil}
	err := Run(data, in, out)
	if err != nil {
		t.Fatalf("Problem running data: %v", err)
	}
	assertOutput(t, expected, out)
}

func assertOutput(t *testing.T, expected []int, out *slicedest) {
	if !reflect.DeepEqual(expected, out.output) {
		t.Fatalf("Expected output to be %v, but was %v", expected, out.output)
	}
}

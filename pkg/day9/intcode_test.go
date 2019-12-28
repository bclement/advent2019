package day9

import (
	"reflect"
	"testing"
)

func TestQuine(t *testing.T) {
	data := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	var io SliceIO
	comp := NewComputer(data, &io, &io)
	err := comp.Run()
	if err != nil {
		t.Fatalf("Error running program: %v", err)
	}
	if !reflect.DeepEqual(io.output, data) {
		t.Fatalf("Expected %v got %v", data, io.output)
	}
}

func TestLargeNum(t *testing.T) {
	data := []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	var io SliceIO
	comp := NewComputer(data, &io, &io)
	err := comp.Run()
	if err != nil {
		t.Fatalf("Error running program: %v", err)
	}
	expected := []int64{1219070632396864}
	if !reflect.DeepEqual(io.output, expected) {
		t.Fatalf("Expected %v got %v", expected, io.output)
	}
}

func TestLargeWrite(t *testing.T) {
	data := []int64{104, 1125899906842624, 99}
	var io SliceIO
	comp := NewComputer(data, &io, &io)
	err := comp.Run()
	if err != nil {
		t.Fatalf("Error running program: %v", err)
	}
	expected := []int64{1125899906842624}
	if !reflect.DeepEqual(io.output, expected) {
		t.Fatalf("Expected %v got %v", expected, io.output)
	}
}

func TestDiagnostic(t *testing.T) {
	assertFull(t, 1, 2932210790)
}

func TestBoost(t *testing.T) {
	assertFull(t, 2, 73144)
}

func assertFull(t *testing.T, input, expected int64) {
	data, err := ReadDataFile("../../resources/day9/input.txt")
	if err != nil {
		t.Fatalf("Problem reading input: %v", err)
	}
	io := NewSliceIO([]int64{input})
	comp := NewComputer(data, io, io)
	err = comp.Run()
	if err != nil {
		t.Fatalf("Error running program: %v", err)
	}
	if !reflect.DeepEqual(io.output[0], expected) {
		t.Fatalf("Expected %v got %v", expected, io.output)
	}
}

func TestInOut(t *testing.T) {
	in := []int64{42}
	data := []int64{3, 0, 4, 0, 99}
	assertRunData(t, data, in, in)
}

func TestPosEq(t *testing.T) {
	in := []int64{8}
	data := []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	expected := []int64{1}
	assertRunData(t, data, in, expected)
}

func TestPosNEq(t *testing.T) {
	data := []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int64{42}
	expected := []int64{0}
	assertRunData(t, data, in, expected)
}

func TestPosLT(t *testing.T) {
	data := []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int64{7}
	expected := []int64{1}
	assertRunData(t, data, in, expected)
}

func TestPosGTE(t *testing.T) {
	data := []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	in := []int64{9}
	expected := []int64{0}
	assertRunData(t, data, in, expected)
}

func TestImmEq(t *testing.T) {
	data := []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	in := []int64{8}
	expected := []int64{1}
	assertRunData(t, data, in, expected)
}

func TestPJump(t *testing.T) {
	data := []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	in := []int64{0}
	expected := []int64{0}
	assertRunData(t, data, in, expected)
}

func TestPJumpNE(t *testing.T) {
	data := []int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	in := []int64{42}
	expected := []int64{1}
	assertRunData(t, data, in, expected)
}

func assertRunData(t *testing.T, data, input, expected []int64) {
	io := NewSliceIO(input)
	comp := NewComputer(data, io, io)
	err := comp.Run()
	if err != nil {
		t.Fatalf("Problem running data: %v", err)
	}
	assertOutput(t, expected, io)
}

func assertOutput(t *testing.T, expected []int64, out *SliceIO) {
	if !reflect.DeepEqual(expected, out.output) {
		t.Fatalf("Expected output to be %v, but was %v", expected, out.output)
	}
}

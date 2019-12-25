package day8

import (
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {
	img, err := openImage()
	if err != nil {
		t.Fatalf("problem parsing input: %v", err)
	}
	checksum := img.getChecksum()
	expected := 1224
	if checksum != expected {
		t.Errorf("Expected checksum %v got %v", expected, checksum)
	}
}

func TestRender(t *testing.T) {
	img, err := openImage()
	if err != nil {
		t.Fatalf("problem parsing input: %v", err)
	}
	result := img.render()
	if result != "" {
		t.Errorf("\n%v\n", result)
	}
}

func openImage() (*image, error) {
	f, err := os.Open("../../resources/day8/input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readImage(f, 25, 6)
}

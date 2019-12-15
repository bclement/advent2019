package day1

import (
	"bufio"
	"io"
	"strconv"
)

// IntLineReader scans files with an int on each line
type IntLineReader struct {
	scanner *bufio.Scanner
}

// NewIntLineReader creates an int scanner from a reader
func NewIntLineReader(in io.Reader) IntLineReader {
	return IntLineReader{bufio.NewScanner(in)}
}

// HasNext returns true if there is another line false otherwise
func (r IntLineReader) HasNext() bool {
	return r.scanner.Scan()
}

// Next returns the next int from the reader or an error
func (r IntLineReader) Next() (int, error) {
	return strconv.Atoi(r.scanner.Text())
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateModuleFuel(mass int) int {
	fuel := calculateFuel(mass)
	fuelMass := fuel
	for fuelMass > 0 {
		fuelMass = calculateFuel(fuelMass)
		if fuelMass > 0 {
			fuel += fuelMass
		}
	}
	return fuel
}

func calculateTotalFuel(in io.Reader, addFuelFuel bool) (int, error) {
	reader := NewIntLineReader(in)
	total := 0
	for reader.HasNext() {
		mass, err := reader.Next()
		if err != nil {
			return -1, err
		}
		if addFuelFuel {
			total += calculateModuleFuel(mass)
		} else {
			total += calculateFuel(mass)
		}
	}
	return total, nil
}

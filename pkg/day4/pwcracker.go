package day4

import (
	"strconv"
)

type state int

const (
	start = iota
	seenTwo
	longRun
	hasDouble
)

func getPossibleCountInRange(lower, upper int) int {
	count := 0
	for i := lower; i < upper; i++ {
		pw := strconv.Itoa(i)
		if possiblePassword(pw) {
			count++
		}
	}
	return count
}

func possiblePassword(pw string) bool {
	if len(pw) != 6 {
		return false
	}
	s := start
	lastDigit := -1
	for i := 0; i < len(pw); i++ {
		digit, err := strconv.Atoi(pw[i : i+1])
		if err != nil {
			return false
		}
		if digit < lastDigit {
			return false
		}
		switch s {
		case start:
			if digit == lastDigit {
				s = seenTwo
			}
		case seenTwo:
			if digit == lastDigit {
				s = longRun
			} else {
				s = hasDouble
			}
		case longRun:
			if digit != lastDigit {
				s = start
			}
		}
		lastDigit = digit
	}
	return s == hasDouble || s == seenTwo
}

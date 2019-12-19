package day4

import "testing"

func TestInput(t *testing.T) {
	assertPw(t, "111111", false)
	assertPw(t, "223450", false)
	assertPw(t, "123789", false)
	assertPw(t, "112233", true)
	assertPw(t, "123444", false)
	assertPw(t, "111122", true)
	assertPw(t, "223333", true)
	assertPw(t, "456777", false)
	assertPw(t, "446777", true)
	assertPw(t, "444567", false)
}

func assertPw(t *testing.T, pw string, expected bool) {
	result := possiblePassword(pw)
	if result != expected {
		t.Errorf("Expected %v for %v, got %v", expected, pw, result)
	}
}

func TestCountPossiblePWs(t *testing.T) {
	result := getPossibleCountInRange(165432, 707912)
	if result != 1163 {
		t.Errorf("Unexpected result: %v", result)
	}
}

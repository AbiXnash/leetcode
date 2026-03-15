package main

import (
	"testing"
)

func TestRomanToInt1(t *testing.T) {
	result := RomanToInt("III")
	if result != 3 {
		t.Errorf("Failed 1")
	}
}

func TestRomanToInt2(t *testing.T) {
	result := RomanToInt("LVIII")
	if result != 58 {
		t.Errorf("Failed 2")
	}
}

func TestRomanToInt3(t *testing.T) {
	result := RomanToInt("MCMXCIV")
	if result != 1994 {
		t.Errorf("Failed 3")
	}
}

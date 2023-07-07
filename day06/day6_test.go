package day06

import (
	"testing"
)

func TestDay6P1(t *testing.T) {
	input := []string{"abcabcabcd"}

	actual := day6PartOne(input)
	expected := 10

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay6P2(t *testing.T) {
	input := []string{"aaabbbabcdefghijklmn"}

	actual := day6PartTwo(input)
	expected := 20

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

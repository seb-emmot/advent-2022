package main

import "testing"

func AreEqual(a any, b any) bool {
	return a == b
}

func TestGetValue(t *testing.T) {
	input := rune('A')
	expected := 27
	actual := getValue(input)

	if !AreEqual(expected, actual) {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestGetDuplicate(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
	expected := "p"

	actual := getDuplicate(input)

	if !AreEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestFindCommon(t *testing.T) {
	input := []string{"abc", "defb", "ghib"}
	expected := 'b'

	actual := findCommon(input)

	if !AreEqual(expected, actual) {
		t.Errorf("Expected %c, got %c", expected, actual)
	}
}

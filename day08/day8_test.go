package day08

import (
	"advent/utils"
	"testing"
)

func TestDay8PartOne(t *testing.T) {
	input := utils.ReadFile("../test_input/day8.txt")

	actual := day8PartOne(input)
	expected := 21

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay8PartOneFull(t *testing.T) {
	input := utils.ReadFile("../inputs/day8.txt")

	actual := day8PartOne(input)
	expected := 1787

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestNewHeightMap(t *testing.T) {
	hm := NewHeightMap([]string{"123", "456", "789"})

	actual := hm.Circumferance()
	expected := 8

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestComputeVisible(t *testing.T) {
	hm := NewHeightMap([]string{"000", "010", "000"})

	actual := len(hm.ComputeVisible())
	expected := 1

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestComputeVisibleVertical(t *testing.T) {
	hm := NewHeightMap([]string{"000", "212", "000"})

	actual := len(hm.ComputeVisible())
	expected := 1

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestScenicScore(t *testing.T) {
	input := utils.ReadFile("../test_input/day8.txt")
	hm := NewHeightMap(input)

	actual := hm.ScenicScore(2, 1)
	expected := 4

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestScenicScores(t *testing.T) {
	input := utils.ReadFile("../test_input/day8.txt")
	hm := NewHeightMap(input)

	scores := hm.ScenicScores()
	actual := utils.MaxArray(scores)
	expected := 8

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestScenicScoresFull(t *testing.T) {
	input := utils.ReadFile("../inputs/day8.txt")
	hm := NewHeightMap(input)

	scores := hm.ScenicScores()
	actual := utils.MaxArray(scores)
	expected := 440640

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

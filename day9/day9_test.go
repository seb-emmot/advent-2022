package day9

import (
	"advent/utils"
	"testing"
)

func TestParseMoves(t *testing.T) {
	input := []string{"R 2", "U 2", "L 1", "D 3"}

	moves := parseMoves(input)

	expected := []Move{Right, Right, Up, Up, Left, Down, Down, Down}

	for i, move := range moves {
		if move != expected[i] {
			t.Errorf("Failure")
		}
	}
}

func TestPartOne(t *testing.T) {
	b := MakeRopeGame(2)

	moves := parseMoves(utils.ReadFile("../inputs/day9.txt"))

	b.DoMoves(moves)

	actual := len(utils.Set(b.TailVisited))
	expected := 5710

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	b := MakeRopeGame(10)

	moves := parseMoves(utils.ReadFile("../inputs/day9.txt"))

	b.DoMoves(moves)

	actual := len(utils.Set(b.TailVisited))
	expected := 2259

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

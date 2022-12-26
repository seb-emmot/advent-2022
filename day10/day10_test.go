package day10

import (
	"advent/utils"
	"testing"
)

func TestParseInstructions(t *testing.T) {
	input := []string{
		"addx 13\n",
		"noop\n",
		"addx -5\n",
	}

	instr := parseInstructions(input)

	print(instr)

}

func TestRunProgramPartOne(t *testing.T) {
	input := utils.ReadFile("../inputs/day10.txt")

	instr := parseInstructions(input)
	c := NewComputer()
	actual := c.RunProgram(instr)
	expected := 15020

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestRunProgramPartTwo(t *testing.T) {
	input := utils.ReadFile("../inputs/day10.txt")

	instr := parseInstructions(input)
	c := NewComputer()
	c.RunProgram(instr)

	screen := renderCRT(c.History)

	checkSum := 0
	for y := 0; y < len(screen); y++ {
		for x := 0; x < len(screen[y]); x++ {
			val := screen[y][x]
			checkSum += int(val[0]) * (x + 1) * (y + 1)
		}
	}

	expected := 721115

	if checkSum != expected {
		t.Errorf("Expected %d, got %d", expected, checkSum)
	}
}

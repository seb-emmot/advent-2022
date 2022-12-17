package main

import (
	"day1/utils"
	"strings"
)

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

type Result int

const (
	Lost Result = iota
	Draw
	Win
)

func day2() {
	input := readFile("inputs/day2.txt")

	var r1Scores []int
	var r2Scores []int

	for numRound, moves := range input {

		// Stragey 1 appaoach
		moves := strings.Split(moves, " ")
		elfMove := parseMove(moves[0])
		myMove := parseMove(moves[1])

		r1Res := evalRound(myMove, elfMove)

		r1Score := computeMoveScore(myMove) + computeRoundScore(r1Res)
		r1Scores = append(r1Scores, r1Score)

		// Strategy 2 approach
		desiredOutcome := getOutcome(moves[1])
		desiredMove := getMoveForOutcome(desiredOutcome, elfMove)

		r2Res := evalRound(desiredMove, elfMove)

		r2Score := computeMoveScore(desiredMove) + computeRoundScore(r2Res)
		r2Scores = append(r2Scores, r2Score)

		println(
			"round #", numRound,
			"moves ", elfMove, myMove,
			"score", r1Score)
	}

	sumR1 := utils.SumArray(r1Scores)
	sumR2 := utils.SumArray(r2Scores)

	println("total score R1", sumR1)
	println("total score R2", sumR2)
}

func computeMoveScore(move Move) int {
	switch move {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic(-1)
}

func computeRoundScore(res Result) int {
	switch res {
	case Lost:
		return 0
	case Draw:
		return 3
	case Win:
		return 6
	}
	panic(-1)
}

func parseMove(move string) Move {
	switch move {
	case "X":
		return Rock
	case "A":
		return Rock
	case "Y":
		return Paper
	case "B":
		return Paper
	case "Z":
		return Scissors
	case "C":
		return Scissors
	}

	panic(-1)
}

func getOutcome(outcome string) Result {
	switch outcome {
	case "X":
		return Lost
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	panic(-1)
}

func getMoveForOutcome(outcome Result, opponentMove Move) Move {
	switch outcome {
	case Lost:
		switch opponentMove {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	case Draw:
		return opponentMove
	case Win:
		switch opponentMove {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}

	panic(-1)
}

func evalRound(me, elf Move) Result {
	if me == elf {
		return Draw
	}

	if me == Rock && elf == Paper {
		return Lost
	}

	if me == Rock && elf == Scissors {
		return Win
	}

	if me == Paper && elf == Rock {
		return Win
	}

	if me == Paper && elf == Scissors {
		return Lost
	}

	if me == Scissors && elf == Rock {
		return Lost
	}

	if me == Scissors && elf == Paper {
		return Win
	}

	return Lost
}

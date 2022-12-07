package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []string {
	readFile, err := os.Open(path)
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	return lines
}

func day1() {
	inputLines := readFile("inputs/day1.txt")

	elfCount := 0
	elfSum := 0

	var elfSums []int

	for _, x := range inputLines {
		if x == "" {
			elfSums = append(elfSums, elfSum)
			elfSum = 0
			elfCount++
		} else {
			n, e := strconv.Atoi(x)
			check(e)

			elfSum += n
		}
	}

	// extract the 3 largest values.
	sort.Ints(elfSums)
	length := len(elfSums)
	subset := elfSums[length-3 : length]

	fmt.Println("top 3 elves", subset)

	// sum the subset.
	sum := 0
	for _, v := range subset {
		sum += v
	}

	fmt.Println("total", sum)
}

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

	sumR1 := sumArray(r1Scores)
	sumR2 := sumArray(r2Scores)

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

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func main() {
	day2()
}

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

type Moves int

const (
	Rock Moves = iota
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

	var scores []int

	for numRound, moves := range input {
		moves := strings.Split(moves, " ")
		elfMove := parseMove(moves[0])
		myMove := parseMove(moves[1])

		score := 0

		switch myMove {
		case Rock:
			score += 1
		case Paper:
			score += 2
		case Scissors:
			score += 3
		}

		res := evalRound(myMove, elfMove)

		switch res {
		case Lost:
			score += 0
		case Draw:
			score += 3
		case Win:
			score += 6
		}

		scores = append(scores, score)

		println(
			"round #", numRound,
			"moves ", elfMove, myMove,
			"score", score)
	}

	sum := sumArray(scores)

	println("total score", sum)
}

func parseMove(move string) Moves {
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

	return Rock
}

func evalRound(me, elf Moves) Result {
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

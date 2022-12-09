package main

import "strings"

func day3() {
	input := readFile("inputs/day3.txt")

	score := 0

	for _, line := range input {
		dupe := getDuplicate(line)
		prio := getValue(rune(dupe[0]))

		score += prio
	}

	println(score)
}

func getDuplicate(items string) string {
	len := len(items)
	if len%2 > 0 {
		panic("items must be even")
	}

	firstCompartment := items[0 : len/2]
	secondCompartment := items[len/2 : len]

	for _, char := range firstCompartment {
		if strings.ContainsRune(secondCompartment, char) {
			return string(char)
		}
	}
	panic("no duplicate found")
}

// getValue returns the value of a character
// according to the rules of aoc day3.
func getValue(char rune) int {
	if char > 90 {
		// lowercase
		return int(char) - 96
	} else {
		// uppercase
		return int(char) - 38
	}
}

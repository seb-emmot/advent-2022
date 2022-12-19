package day3

import (
	"advent/utils"
	"strings"
)

func day3() {
	input := utils.ReadFile("inputs/day3.txt")

	score := 0

	for _, line := range input {
		// make sure line is even length.
		utils.AssertDivisibleStr(line, 2)
		firstHalf := line[0 : len(line)/2]
		secHalf := line[len(line)/2:]

		dupe := findCommon([]string{firstHalf, secHalf})
		prio := getValue(dupe)

		score += prio
	}

	bInput := utils.Reshape(input, 3)
	tokenScore := 0

	for _, group := range bInput {
		token := findCommon(group)
		tokenScore += getValue(token)

		println(string(token))
	}

	println("part1 score:", score)
	println("part2 score:", tokenScore)

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

func findCommon(strs []string) rune {
	if len(strs) < 2 {
		panic("need at least 2 strings")
	}

	var unique rune
	for _, char := range strs[0] {
		var isCandidate bool = true
		for _, str := range strs[1:] {
			if !strings.ContainsRune(str, char) {
				isCandidate = false
				break
			}
		}
		if isCandidate {
			unique = char
			break
		}
	}

	return unique
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

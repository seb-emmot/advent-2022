package main

import (
	"advent/day9"
	"advent/utils"
)

func main() {
	input := utils.ReadFile("inputs/day9.txt")
	println("p1", day9.Day9PartOne(input))
	println("p2", day9.Day9PartTwo(input))
}

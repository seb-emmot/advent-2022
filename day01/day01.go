package day01

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

func day1() {
	inputLines := utils.ReadFile("inputs/day1.txt")

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
			utils.Check(e)

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

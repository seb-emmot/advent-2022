package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func main() {
	day1()
}

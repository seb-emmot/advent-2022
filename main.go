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

func day1() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elfCount := 0
	elfSum := 0

	var elfSums []int

	for fileScanner.Scan() {
		x := fileScanner.Text()

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

	readFile.Close()
}

func main() {
	day1()
}

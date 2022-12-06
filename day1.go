package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elfCount := 0
	elfSum := 0

	currentMax := 0

	for fileScanner.Scan() {
		x := fileScanner.Text()

		if x == "" {
			if elfSum > currentMax {
				currentMax = elfSum
			}
			fmt.Println("elf", elfCount, "sum", elfSum)
			elfSum = 0
			elfCount++
		} else {
			n, e := strconv.Atoi(x)
			check(e)

			elfSum += n
		}
	}

	// print max elfSum
	fmt.Println("max", currentMax)

	readFile.Close()
}

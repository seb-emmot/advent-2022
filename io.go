package main

import (
	"bufio"
	"day1/utils"
	"os"
)

func readFile(path string) []string {
	readFile, err := os.Open(path)
	utils.Check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	return lines
}

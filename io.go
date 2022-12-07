package main

import (
	"bufio"
	"os"
)

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

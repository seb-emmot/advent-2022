package main

func containsDuplicate(str string) bool {
	seen := make(map[string]bool)
	for _, c := range str {
		if seen[string(c)] {
			return true
		}
		seen[string(c)] = true
	}
	return false
}

func findUniqueSequencePos(input string, numReq int) int {
	for i := numReq; i <= len(input); i++ {
		if !containsDuplicate(input[i-numReq : i]) {
			return i
		}
	}
	panic("no solution found")
}

func day6PartOne(input []string) int {
	stream := input[0]
	res := findUniqueSequencePos(stream, 4)

	return res
}

func day6PartTwo(input []string) int {
	stream := input[0]
	res := findUniqueSequencePos(stream, 14)

	return res
}

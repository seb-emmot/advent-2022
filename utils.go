package main

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
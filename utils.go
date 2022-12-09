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

func printList(list []string) {
	for _, v := range list {
		println(v)
	}
}

func reshape(items []string, size int) [][]string {
	var reshaped [][]string
	assertDivisibleArr(items, size)

	for i := 0; i < len(items); i += size {
		reshaped = append(reshaped, items[i:i+size])
	}

	return reshaped
}

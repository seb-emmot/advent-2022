package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	stacks map[int]*Stack[string]
}

func (b *Board) Add(stack int, item string) {
	if b.stacks[stack] == nil {
		b.stacks[stack] = new(Stack[string])
	}
	b.stacks[stack].Push(item)
}

func (b *Board) Take(stack int) string {
	if b.stacks[stack] == nil {
		panic("stack is empty!")
	}
	return b.stacks[stack].Pop()
}

func (b *Board) Print() {
	for i := 1; i <= len(b.stacks); i++ {
		fmt.Println(i, b.stacks[i].s)
	}
}

func (b *Board) Top() string {
	str := ""
	for i := 1; i <= len(b.stacks); i++ {
		stack := b.stacks[i].s
		str += stack[len(stack)-1]
	}
	return str
}

func splitBy(input []string, sep string) ([]string, []string) {
	for i, line := range input {
		if line == sep {
			return input[:i], input[i+1:]
		}
	}
	panic("No separator found")
}

func tryParseInt(a rune, out *int) bool {
	res, err := strconv.Atoi(string(a))
	if err != nil {
		return false
	} else {
		*out = res
		return true
	}
}

func tryParseChar(a rune, out *string) bool {
	if a >= 65 && a <= 90 {
		*out = string(a)
		return true
	}
	return false
}

func day5() string {
	input := readFile("inputs/day5.txt")

	initial, moves := splitBy(input, "")

	b := new(Board)
	b.stacks = make(map[int]*Stack[string])

	for row := len(initial) - 1; row >= 0; row-- {
		for col, c := range initial[row] {
			res := ""
			if tryParseChar(c, &res) {
				pos := (col + 3) / 4
				b.Add(pos, res)
			}
		}
	}
	b.Print()
	fmt.Println()

	// parse the moves
	for _, move := range moves {

		splt := strings.Split(move, " ")

		num, e1 := strconv.Atoi(splt[1])
		check(e1)
		from, e2 := strconv.Atoi(splt[3])
		check(e2)
		to, e3 := strconv.Atoi(splt[5])
		check(e3)

		for i := 0; i < num; i++ {
			res := b.Take(from)
			b.Add(to, res)
		}
	}

	b.Print()

	return b.Top()
}

//SHQWSRBD

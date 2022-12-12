package main

import "testing"

func TestDay5(t *testing.T) {
	r1, r2 := day5()

	if r1 != "SHQWSRBDL" {
		panic("Expected SHQWSRBDL, got " + r1)
	}

	if r2 != "CDTQZHBRS" {
		panic("Expected CDTQZHBRS, got " + r2)
	}
}

func TestBoard(t *testing.T) {
	b := new(Board)
	b.stacks = make(map[int]*Stack[string])

	b.Add(1, "a")
	b.Add(1, "b")

	b.Add(2, "c")
	res := b.Take(1)
	b.Add(2, res)

	ret := b.Take(2)
	if ret != "b" {
		t.Errorf("Expected b, got %s", ret)
	}
}

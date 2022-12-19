package day4

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Span struct {
	Min int
	Max int
}

func (first Span) Contains(sub Span) bool {
	return first.Min <= sub.Min && first.Max >= sub.Max
}

func (first Span) Overlaps(sub Span) bool {
	return first.Min <= sub.Max && first.Max >= sub.Min
}

func buildSpans(input []string) []Span {
	var spans []Span
	for _, v := range input {
		x := strings.Split(v, ",")
		for _, q := range x {
			a := strings.Split(q, "-")
			min, e1 := strconv.Atoi(a[0])
			max, e2 := strconv.Atoi(a[1])
			utils.Check(e1)
			utils.Check(e2)

			s := Span{min, max}
			spans = append(spans, s)
		}
	}

	return spans
}

func checkContains(spans []Span) []bool {
	var res []bool
	for i := 0; i < len(spans); i += 2 {
		first := spans[i]
		second := spans[i+1]

		var contains bool = first.Contains(second) || second.Contains(first)

		res = append(res, contains)
	}
	return res
}

func checkOverlaps(spans []Span) []bool {
	var res []bool
	for i := 0; i < len(spans); i += 2 {
		first := spans[i]
		second := spans[i+1]

		var overlaps bool = first.Overlaps(second)

		res = append(res, overlaps)
	}
	return res
}

func day4() {
	input := utils.ReadFile("inputs/day4.txt")

	spans := buildSpans(input)

	res := checkContains(spans)

	p1Count := 0
	for _, v := range res {
		if v {
			p1Count++
		}
	}

	r2 := checkOverlaps(spans)

	p2Count := 0
	for _, v := range r2 {
		if v {
			p2Count++
		}
	}

	fmt.Println("p1: ", p1Count, "p2: ", p2Count)
}

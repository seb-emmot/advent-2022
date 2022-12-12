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

type Queue[T any] struct {
	q []T
}

func (q *Queue[T]) Push(v T) {
	if q.q == nil {
		q.q = make([]T, 0)
	}
	q.q = append(q.q, v)
}

func (q *Queue[T]) Pop() T {
	res := q.q[0]
	q.q = q.q[1:]
	return res
}

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) Push(v T) {
	if s.s == nil {
		s.s = make([]T, 0)
	}
	s.s = append(s.s, v)
}

func (s *Stack[T]) PushMany(v []T) {
	if s.s == nil {
		s.s = make([]T, 0)
	}
	s.s = append(s.s, v...)
}

func (s *Stack[T]) Pop() T {
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

func (q *Stack[T]) PopMany(num int) []T {
	res := q.s[len(q.s)-num:]
	q.s = q.s[:len(q.s)-num]
	return res
}

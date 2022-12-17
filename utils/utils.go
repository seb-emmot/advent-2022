package utils

func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func printList(list []string) {
	for _, v := range list {
		println(v)
	}
}

func Reshape(items []string, size int) [][]string {
	var reshaped [][]string
	AssertDivisibleArr(items, size)

	for i := 0; i < len(items); i += size {
		reshaped = append(reshaped, items[i:i+size])
	}

	return reshaped
}

func AssertDivisibleArr(items []string, mod int) {
	len := len(items)
	if len%mod > 0 {
		panic("items must be divisible")
	}
}

func AssertDivisibleStr(items string, mod int) {
	len := len(items)
	if len%mod > 0 {
		panic("items must be divisible")
	}
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
	S []T
}

func (s *Stack[T]) Push(v T) {
	if s.S == nil {
		s.S = make([]T, 0)
	}
	s.S = append(s.S, v)
}

func (s *Stack[T]) PushMany(v []T) {
	if s.S == nil {
		s.S = make([]T, 0)
	}
	s.S = append(s.S, v...)
}

func (s *Stack[T]) Pop() T {
	res := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]
	return res
}

func (q *Stack[T]) PopMany(num int) []T {
	res := q.S[len(q.S)-num:]
	q.S = q.S[:len(q.S)-num]
	return res
}

type Point struct {
	X int
	Y int
}

func Set[K Point](input []K) []K {
	m := make(map[K]int)

	for _, a := range input {
		_, exist := m[a]
		if !exist {
			m[a] = 1
		}
	}

	keys := []K{}
	i := 0
	for k := range m {
		keys = append(keys, k)
		i++
	}

	return keys
}

func MaxArray(a []int) int {
	max := a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

package slicer

import (
	"math/rand"
	"testing"
)

type testCase struct {
	operand  int
	sorted   []int
	expected []int
}

// go test ./pkg/slicer
func TestInsert(t *testing.T) {
	cases := []testCase{
		{10, []int{}, []int{10}},
		{10, []int{10}, []int{10, 10}},
		{11, []int{10, 11}, []int{10, 11, 11}},
		{11, []int{3, 4}, []int{3, 4, 11}},
		{1, []int{2, 3}, []int{1, 2, 3}},
		{1, []int{1, 2, 4}, []int{1, 1, 2, 4}},
		{0, []int{1, 500}, []int{0, 1, 500}},
		{0, []int{0, 500}, []int{0, 0, 500}},
	}

	for i, c := range cases {
		actual := Insert(c.operand, c.sorted)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Insert(%d): expected %d, actual %d", i, c.operand, c.expected, actual)
		}
	}
}

func TestDelete(t *testing.T) {
	cases := []testCase{
		{10, []int{}, []int{}},
		{10, []int{10}, []int{}},
		{11, []int{10, 11}, []int{10}},
		{11, []int{10, 11, 11, 11, 11}, []int{10}},
		{0, []int{0, 3, 4}, []int{3, 4}},
		{1, []int{2, 3, 1}, []int{2, 3}},
		{1, []int{2, 3, 1, 1}, []int{2, 3}},
		{1, []int{1, 2, 3}, []int{2, 3}},
		{1, []int{1, 1, 2, 3}, []int{2, 3}},
	}

	for i, c := range cases {
		actual := Delete(c.operand, c.sorted)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Delete(%d): expected %d, actual %d", i, c.operand, c.expected, actual)
		}
	}
}

// go test -bench=. -benchmem ./pkg/slicer
var result []int
var n = 1000

// Кейс с превыделением памяти для слайса и append (чуть быстрее чем c copy)
func BenchmarkInsert1_5(b *testing.B)   { benchmarkInsert(5, b) }
func BenchmarkInsert1_50(b *testing.B)  { benchmarkInsert(50, b) }
func BenchmarkInsert1_100(b *testing.B) { benchmarkInsert(100, b) }

// Кейс с превыделением памяти для слайса и copy
func BenchmarkInsert2_5(b *testing.B)   { benchmarkInsert2(5, b) }
func BenchmarkInsert2_50(b *testing.B)  { benchmarkInsert2(50, b) }
func BenchmarkInsert2_100(b *testing.B) { benchmarkInsert2(100, b) }

func BenchmarkDelete5(b *testing.B)   { benchmarkDelete(5, b) }
func BenchmarkDelete50(b *testing.B)  { benchmarkDelete(50, b) }
func BenchmarkDelete100(b *testing.B) { benchmarkDelete(100, b) }

func benchmarkInsert(size int, b *testing.B) {
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	var sorted = makeSortedSlice(size, n)

	for i := 0; i < b.N; i++ {
		r = Insert(x, sorted)
	}
	result = r
}

func benchmarkInsert2(size int, b *testing.B) {
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	var sorted = makeSortedSlice(size, n)

	for i := 0; i < b.N; i++ {
		r = Insert2(x, sorted)
	}
	result = r
}

func benchmarkDelete(size int, b *testing.B) {
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	var sorted = makeSortedSlice(size, n)

	for i := 0; i < b.N; i++ {
		r = Delete(x, sorted)
	}
	result = r
}

func isEqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func makeSortedSlice(size int, n int) []int {
	var s = make([]int, 0, size)
	n = rand.Intn(n)

	for i := 0; i < size; i++ {
		s = append(s, n+i)
	}

	return s
}

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
func TestManagerInsert(t *testing.T) {
	var manager SliceManager
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
		actual := manager.Insert(c.operand, c.sorted)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Insert(%d): expected %d, actual %d", i, c.operand, c.expected, actual)
		}
	}
}

func TestManagerDelete(t *testing.T) {
	var manager SliceManager
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
		actual := manager.Delete(c.operand, c.sorted)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Delete(%d): expected %d, actual %d", i, c.operand, c.expected, actual)
		}
	}
}

// go test -bench=. -benchmem ./pkg/slicer
var result []int
var n = 1000

// Кейс с превыделением памяти для слайса и append (чуть быстрее чем c copy)
func BenchmarkManagerInsert5(b *testing.B)   { benchmarkManagerInsert(5, b) }
func BenchmarkManagerInsert50(b *testing.B)  { benchmarkManagerInsert(50, b) }
func BenchmarkManagerInsert100(b *testing.B) { benchmarkManagerInsert(100, b) }

func BenchmarkManagerDelete5(b *testing.B)   { benchmarkManagerDelete(5, b) }
func BenchmarkManagerDelete50(b *testing.B)  { benchmarkManagerDelete(50, b) }
func BenchmarkManagerDelete100(b *testing.B) { benchmarkManagerDelete(100, b) }

func benchmarkManagerInsert(size int, b *testing.B) {
	var manager SliceManager
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	var sorted = makeSortedSlice(size, n)

	for i := 0; i < b.N; i++ {
		r = manager.Insert(x, sorted)
	}
	result = r
}

func benchmarkManagerDelete(size int, b *testing.B) {
	var manager SliceManager
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	var sorted = makeSortedSlice(size, n)

	for i := 0; i < b.N; i++ {
		r = manager.Delete(x, sorted)
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

package slicer

import (
	"math/rand"
	"testing"
)

type testCase struct {
	value    int
	items    []int
	expected []int
}

// go test ./pkg/slicer
func TestManagerInsert(t *testing.T) {
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
		manager := SliceManager{c.items}
		actual := manager.Insert(c.value)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Insert(%d): expected %d, actual %d", i, c.value, c.expected, actual)
		}
	}
}

func TestManagerDelete(t *testing.T) {
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
		manager := SliceManager{c.items}
		actual := manager.Delete(c.value)
		if !isEqualSlices(actual, c.expected) {
			t.Errorf("Usecase [%d]. Delete(%d): expected %d, actual %d", i, c.value, c.expected, actual)
		}
	}
}

func TestManagerIsEqual(t *testing.T) {
	cases := []struct {
		slicer1  SliceManager
		slicer2  SliceManager
		expected bool
	}{
		{SliceManager{}, SliceManager{}, true},
		{SliceManager{[]int{0}}, SliceManager{[]int{0}}, true},
		{SliceManager{[]int{1}}, SliceManager{}, false},
		{SliceManager{[]int{25}}, SliceManager{[]int{25}}, true},
		{SliceManager{[]int{10, 11, 11, 11, 11}}, SliceManager{[]int{10, 11, 11, 11, 11}}, true},
		{SliceManager{[]int{10, 11, 11, 11}}, SliceManager{[]int{10, 11, 11, 11, 11}}, false},
	}

	for i, c := range cases {
		if c.slicer1.isEqual(c.slicer2) != c.expected {
			t.Errorf("Usecase [%d]. isEqual(...): expected %v, actual %v", i, c.expected, c.slicer1.isEqual(c.slicer2))
		}
	}
}

func TestManagerGetMax(t *testing.T) {
	manager := SliceManager{}
	sorted := [5]int{10, 250, 251, 251, 255}
	for i := 0; i < 5; i++ {
		manager.Insert(sorted[i])
		max, _ := manager.getMax()
		if max != sorted[i] {
			t.Errorf("getMax test usecase [%d]. insert(%d): actual %d", i, sorted[i], max)
		}
	}

	manager.Insert(1)
	max, _ := manager.getMax()
	if max != 255 {
		t.Errorf("getMax test. insert low value %d and getMax %d expected %d", 1, max, 255)
	}

	s := SliceManager{}
	_, err := s.getMax()

	if err == nil {
		t.Errorf("Invalid getMax() err value for emty manager")
	}

	if err != nil && err.Error() != "the list is empty" {
		t.Errorf("Invalid getMax() err message")
	}
}

func TestManagerGetMin(t *testing.T) {
	manager := SliceManager{}
	sorted := [5]int{100, 90, 80, 70, 60}

	for i := 0; i < 5; i++ {
		manager.Insert(sorted[i])
		max, _ := manager.getMin()
		if max != sorted[i] {
			t.Errorf("getMin test usecase [%d]. insert(%d): actual %d", i, sorted[i], max)
		}
	}

	manager.Insert(10000)
	min, _ := manager.getMin()
	if min != 60 {
		t.Errorf("getMin test. insert low value %d and getMin %d expected %d", 1, min, 60)
	}

	s := SliceManager{}
	_, err := s.getMin()

	if err == nil {
		t.Errorf("Invalid getMin() err value for emty manager")
	}

	if err != nil && err.Error() != "the list is empty" {
		t.Errorf("Invalid getMin() err message")
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
	var r []int
	items := makeSortedSlice(size)
	tmp := make([]int, len(items))
	rand.Seed(1)
	var x = rand.Intn(n)

	for i := 0; i < b.N; i++ {
		copy(tmp, items)
		manager := SliceManager{tmp}
		r = manager.Insert(x)
	}

	result = r
}

func benchmarkManagerDelete(size int, b *testing.B) {
	var r []int
	rand.Seed(1)
	var x = rand.Intn(n)
	manager := SliceManager{makeSortedSlice(size)}

	for i := 0; i < b.N; i++ {
		r = manager.Delete(x)
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

func makeSortedSlice(size int) []int {
	var s = make([]int, 0, size)

	for i := 0; i < size; i++ {
		s = append(s, i)
	}

	return s
}

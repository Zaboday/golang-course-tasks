package linkedlist

import (
	"testing"
)

type testCase struct {
	newValue      int
	valuesInList  []int
	expectedChain string
	expectedLen   int
}

// go test ./pkg/linkedlist
func TestLinkedListInsert(t *testing.T) {
	cases := []testCase{
		{10, []int{}, "10", 1},
		{10, []int{10}, "10 -> 10", 2},
		{11, []int{10, 11}, "10 -> 11 -> 11", 3},
		{11, []int{3, 4}, "3 -> 4 -> 11", 3},
		{1, []int{2, 3}, "1 -> 2 -> 3", 3},
		{1, []int{1, 2, 4}, "1 -> 1 -> 2 -> 4", 4},
		{0, []int{1, 500}, "0 -> 1 -> 500", 3},
		{0, []int{0, 500}, "0 -> 0 -> 500", 3},
	}

	for i, c := range cases {
		var list List

		for j := 0; j < len(c.valuesInList); j++ {
			list.Insert(c.valuesInList[j])
		}

		list.Insert(c.newValue)

		actual := list.DisplayChain()
		if actual != c.expectedChain {
			t.Errorf("Usecase [%d]. Insert(%d): expectedChain %s, actual %s", i, c.newValue, c.expectedChain, actual)
		}

		if list.len != c.expectedLen {
			t.Errorf("Usecase [%d]. Insert(%d): expectedlen %d, actual %d", i, c.newValue, c.expectedLen, list.len)
		}
	}

}

func TestLinkedListDelete(t *testing.T) {
	cases := []testCase{
		{10, []int{}, "", 0},
		{10, []int{10}, "", 0},
		{11, []int{10, 11}, "10", 1},
		{11, []int{10, 11, 11, 11, 11}, "10", 1},
		{0, []int{0, 3, 4}, "3 -> 4", 2},
		{1, []int{2, 3, 1}, "2 -> 3", 2},
		{1, []int{2, 3, 1, 1}, "2 -> 3", 2},
		{1, []int{1, 2, 3}, "2 -> 3", 2},
		{1, []int{1, 1, 2, 3}, "2 -> 3", 2},
	}

	for i, c := range cases {
		var list List

		for j := 0; j < len(c.valuesInList); j++ {
			list.Insert(c.valuesInList[j])
		}

		list.Delete(c.newValue)

		actual := list.DisplayChain()
		if actual != c.expectedChain {
			t.Errorf("Usecase [%d]. Delete(%d): expectedChain %s, actual %s", i, c.newValue, c.expectedChain, actual)
		}

		if list.len != c.expectedLen {
			t.Errorf("Usecase [%d]. Delete(%d): expectedlen %d, actual %d", i, c.newValue, c.expectedLen, list.len)
		}
	}
}

func TestLinkedListGetMax(t *testing.T) {
	var list List
	sorted := [5]int{10, 250, 251, 251, 255}
	for i := 0; i < 5; i++ {
		list.Insert(sorted[i])
		max, _ := list.getMax()
		if max != sorted[i] {
			t.Errorf("getMax test usecase [%d]. insert(%d): actual %d", i, sorted[i], max)
		}
	}

	list.Insert(1)
	max, _ := list.getMax()
	if max != 255 {
		t.Errorf("getMax test. insert low value %d and getMax %d expected %d", 1, max, 255)
	}
}

func TestLinkedListGetMin(t *testing.T) {
	var list List
	sorted := [5]int{100, 90, 80, 70, 60}

	for i := 0; i < 5; i++ {
		list.Insert(sorted[i])
		max, _ := list.getMin()
		if max != sorted[i] {
			t.Errorf("getMin test usecase [%d]. insert(%d): actual %d", i, sorted[i], max)
		}
	}

	list.Insert(10000)
	min, _ := list.getMin()
	if min != 60 {
		t.Errorf("getMax test. insert low value %d and getMax %d expected %d", 1, min, 255)
	}
}

func TestLinkedListGetMaxError(t *testing.T) {
	var list List

	max, err := list.getMax()

	if max != 0 {
		t.Errorf("Invalid max value for emty list")
	}

	if err == nil {
		t.Errorf("Invalid err value for emty list")
	}

	if err.Error() != "the list is empty" {

	}
}

func TestLinkedListGetMinError(t *testing.T) {
	var list List

	max, err := list.getMin()

	if max != 0 {
		t.Errorf("Invalid min value for emty list")
	}

	if err == nil {
		t.Errorf("Invalid err value for emty list")
	}

	if err.Error() != "the list is empty" {

	}
}

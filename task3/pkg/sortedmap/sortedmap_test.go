package sortedmap

import (
	"sync"
	"testing"
)

// go test ./pkg/sortedmap

func TestSortedMap_AddItem(t *testing.T) {
	cases := []struct {
		items    []string
		expected map[string]int
	}{
		{[]string{"f", "b"}, map[string]int{"f": 1, "b": 1}},
		{[]string{"f", "f", "g"}, map[string]int{"f": 2, "g": 1}},
		{[]string{}, map[string]int{}},
		{[]string{"f", "F", "g", "g", "1"}, map[string]int{"f": 1, "g": 2, "F": 1, "1": 1}},
	}

	for i, c := range cases {
		var sm = New()

		for _, item := range c.items {
			sm.AddItem(item)
		}

		if !isEqualMaps(sm.items, c.expected) {
			t.Errorf("Usecase [%d]. AddItem(): expected %v, actual %v", i, c.expected, sm.items)
		}
	}
}

func TestSortedMap_AddLineOrder(t *testing.T) {
	cases := []struct {
		items    map[string]int
		expected map[string][2]int
	}{
		{map[string]int{"f": 1, "b": 2}, map[string][2]int{"f": {1, 1}, "b": {2, 1}}},
		{map[string]int{"f": 1, "b": 1, "3": 3}, map[string][2]int{"f": {1, 1}, "b": {1, 1}, "3": {3, 1}}},
	}

	for i, c := range cases {
		var sm = New()

		for item, n := range c.items {
			sm.AddLineOrder(item, n, 1)
		}

		if !isEqualOrders(sm.lineOrder, c.expected) {
			t.Errorf("Usecase [%d]. AddLineOrder(): expected %v, actual %v", i, c.expected, sm.lineOrder)
		}
	}

	var sm = New()

	for _, n := range []int{15, 2, 3, 10, 20, 1, 3} {
		sm.AddLineOrder("foo", 1, n)
	}

	if sm.lineOrder["foo"] != [2]int{1, 1} {
		t.Errorf("AddLineOrder() same value: expected %s, actual %v", "{1, 1}", sm.lineOrder["foo"])
	}
}

func TestSortedMap_AddStopItem(t *testing.T) {
	cases := []struct {
		items    []string
		expected map[string]bool
	}{
		{[]string{"f", "b"}, map[string]bool{"f": true, "b": true}},
		{[]string{"f", "f", "g"}, map[string]bool{"f": true, "g": true}},
		{[]string{}, map[string]bool{}},
		{[]string{"f", "F", "g", "g", "1"}, map[string]bool{"f": true, "g": true, "F": true, "1": true}},
	}

	for i, c := range cases {
		var sm = New()

		for _, item := range c.items {
			sm.AddStopItem(item)
			sm.AddStopItem(item)
		}

		if !isEqualMapsBool(sm.stopItems, c.expected) {
			t.Errorf("Usecase [%d]. AddStopItem(): expected %v, actual %v", i, c.expected, sm.items)
		}
	}
}

func TestSortedMap_Top(t *testing.T) {
	items := map[string]int{"f": 1, "b": 2, "h": 2, "d": 5}
	order := map[string][2]int{"h": {1, 2}, "b": {2, 2}}
	emptyOrder := make(map[string]bool)

	cases := []struct {
		items     map[string]int
		order     map[string][2]int
		stopItems map[string]bool
		topSize   int
		expected  []string
	}{
		{items, make(map[string][2]int), emptyOrder, 1, []string{"d: 5"}},
		{items, order, emptyOrder, 2, []string{"d: 5", "h: 2"}},
		{items, order, emptyOrder, 3, []string{"d: 5", "h: 2", "b: 2"}},
		{items, make(map[string][2]int), map[string]bool{"b": true}, 3, []string{"d: 5", "h: 2", "f: 1"}},
		{items, make(map[string][2]int), map[string]bool{"b": true, "d": true}, 3, []string{"h: 2", "f: 1", ""}},
		{items,
			map[string][2]int{"b": {1, 2}, "h": {2, 1}},
			map[string]bool{"d": true}, 5,
			[]string{"f: 1", "", "", "b: 2", "h: 2"},
		},
	}

	for i, c := range cases {
		sm := SortedMap{c.items, c.order, c.stopItems, new(sync.Mutex)}
		top := sm.Top(c.topSize)

		if !isEqualSlices(c.expected, top) {
			t.Errorf("Usecase [%d]. Top() expected %v, actual %v", i, c.expected, top)
		}
	}
}

func isEqualMaps(a map[string]int, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		t, isOk := b[k]
		if !isOk || t != v {
			return false
		}
	}

	return true
}

func isEqualOrders(a map[string][2]int, b map[string][2]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		t, isOk := b[k]
		if !isOk || t != v {
			return false
		}
	}

	return true
}

func isEqualMapsBool(a map[string]bool, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		t, isOk := b[k]
		if !isOk || t != v {
			return false
		}
	}

	return true
}

// Return true if two slices have same length and same index->values
func isEqualSlices(a []string, b []string) bool {
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

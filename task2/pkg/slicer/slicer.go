package slicer

import (
	"errors"
	"fmt"
)

type SliceManager struct {
}

func (slicer *SliceManager) Insert(x int, sortedSlice []int) []int {
	var tmp = make([]int, 0, len(sortedSlice)+1)

	if len(sortedSlice) == 0 || x < sortedSlice[0] {
		// Добавляем x в начало среза если входной срез пустой или его первый элемент больше чем x
		tmp = append(tmp, x)
		tmp = append(tmp, sortedSlice[0:]...)
	} else {
		for i, elem := range sortedSlice {
			if x <= elem {
				tmp = append(tmp, x)
				tmp = append(tmp, sortedSlice[i:]...)
				break

			}
			tmp = append(tmp, elem)
		}
	}

	// Добавляем в конец среза если x больше последнего элемента
	if len(sortedSlice) != 0 && x > sortedSlice[len(sortedSlice)-1] {
		tmp = append(tmp, x)
	}

	return tmp
}

func (slicer *SliceManager) Delete(x int, anySlice []int) []int {
	var tmp = make([]int, 0, len(anySlice))
	for _, elem := range anySlice {
		if x == elem {
			continue
		}
		tmp = append(tmp, elem)
	}

	return tmp
}

// Linked list implementation

type ListItem struct {
	value      int
	prev, next *ListItem
}

type List struct {
	len  int
	head *ListItem
}

func (L *List) DisplayChain() string {
	var chain string
	item := L.head
	for item != nil {
		if item.next != nil {
			chain = chain + fmt.Sprintf("%d -> ", item.value)
		} else {
			chain = chain + fmt.Sprintf("%d", item.value)
		}
		item = item.next
	}

	return chain
}

func (L *List) Insert(value int) {
	newItem := &ListItem{
		value: value,
	}

	item := L.head
	for item != nil {
		if value <= item.value {

			if item == L.head {
				L.head = newItem
			}

			if item.prev != nil {
				item.prev.next = newItem
				newItem.prev = item.prev
			}

			item.prev = newItem
			newItem.next = item
			break
		}

		if item.next == nil {
			item.next = newItem
			newItem.prev = item
			break
		}

		item = item.next
	}

	if L.head == nil {
		L.head = newItem
	}

	L.len++
}

func (L *List) Delete(value int) {
	item := L.head

	if item == nil {
		return
	}

	for item != nil {
		if item.value == value {

			if item.prev != nil {
				item.prev.next = item.next
			}

			if item.next != nil {
				item.next.prev = item.prev
			}

			if item == L.head {
				L.head = item.next
			}

			L.len--
		}
		item = item.next
	}
}

func (L *List) getMax() (max int, err error) {
	item := L.head
	if item == nil {
		return 0, errors.New("the list is empty")
	}

	max = item.value
	for item != nil {
		if item.value > max {
			max = item.value
		}
		item = item.next
	}
	return max, nil
}

func (L *List) getMin() (min int, err error) {
	item := L.head
	if item == nil {
		return 0, errors.New("the list is empty")
	}

	min = item.value
	for item != nil {
		if item.value < min {
			min = item.value
		}
		item = item.next
	}
	return min, nil
}

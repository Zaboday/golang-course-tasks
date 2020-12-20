package linkedlist

import (
	"errors"
	"fmt"
)

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

func (L *List) Insert(value int) int {
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

	return value
}

func (L *List) Delete(value int) int {
	item := L.head

	if item == nil {
		return 0
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

	return value
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

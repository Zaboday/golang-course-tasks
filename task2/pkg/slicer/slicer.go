package slicer

import "errors"

type SliceManager struct {
	items []int
}

func (slicer *SliceManager) Insert(x int) []int {
	var tmp = make([]int, 0, len(slicer.items)+1)

	if len(slicer.items) == 0 || x < slicer.items[0] {
		// Добавляем x в начало среза если входной срез пустой или его первый элемент больше чем x
		tmp = append(tmp, x)
		tmp = append(tmp, slicer.items[0:]...)
	} else {
		for i, elem := range slicer.items {
			if x <= elem {
				tmp = append(tmp, x)
				tmp = append(tmp, slicer.items[i:]...)
				break

			}
			tmp = append(tmp, elem)
		}
	}

	// Добавляем в конец среза если x больше последнего элемента
	if len(slicer.items) != 0 && x > slicer.items[len(slicer.items)-1] {
		tmp = append(tmp, x)
	}

	slicer.items = tmp

	return tmp
}

func (slicer *SliceManager) Delete(x int) []int {
	var tmp = make([]int, 0, len(slicer.items))
	for _, elem := range slicer.items {
		if x == elem {
			continue
		}
		tmp = append(tmp, elem)
	}

	slicer.items = tmp

	return tmp
}

func (slicer *SliceManager) GetItems() []int {
	tmp := make([]int, len(slicer.items))
	// Возвращаем копию среза
	copy(tmp, slicer.items)

	return tmp
}

func (slicer *SliceManager) getMin() (min int, err error) {
	if len(slicer.items) == 0 {
		return 0, errors.New("the list is empty")
	}

	return slicer.items[0], nil
}

func (slicer *SliceManager) getMax() (max int, err error) {
	if len(slicer.items) == 0 {
		return 0, errors.New("the list is empty")
	}

	return slicer.items[len(slicer.items)-1], nil
}

func (slicer *SliceManager) isEqual(another SliceManager) bool {
	if len(slicer.items) != len(another.items) {
		return false
	}

	for i, value := range another.GetItems() {
		if value != slicer.items[i] {
			return false
		}
	}

	return true
}

package slicer

func Insert(x int, sortedSlice []int) []int {
	var tmp []int
	var added bool

	// Добавляем x в начало среза если входной срез пустой или его первый элемент больше чем x
	if len(sortedSlice) == 0 || x < sortedSlice[0] {
		tmp = append(tmp, x)
		added = true
	}

	for _, elem := range sortedSlice {
		if x <= elem {
			if !added {
				tmp = append(tmp, x)
				added = true
			}
		}
		tmp = append(tmp, elem)
	}

	// Добавляем в конец среза если x больше последнего элемента
	if len(sortedSlice) != 0 && x > sortedSlice[len(sortedSlice)-1] {
		tmp = append(tmp, x)
	}

	return tmp
}

func Delete(x int, anySlice []int) []int {
	var tmp []int
	for _, elem := range anySlice {
		if x == elem {
			continue
		}
		tmp = append(tmp, elem)
	}
	return tmp
}

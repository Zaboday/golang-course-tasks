package slicer

func Insert(x int, sortedSlice []int) []int {
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

func Insert2(x int, sortedSlice []int) []int {
	//TODO: implement insert with copy
	return make([]int, 1)
}

func Delete(x int, anySlice []int) []int {
	var tmp = make([]int, 0, len(anySlice))
	for _, elem := range anySlice {
		if x == elem {
			continue
		}
		tmp = append(tmp, elem)
	}

	return tmp
}

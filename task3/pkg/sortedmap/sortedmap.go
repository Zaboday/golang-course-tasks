package sortedmap

import (
	"fmt"
	"sort"
)

type SortedMap struct {
	items     map[string]int
	order     map[string]int
	stopItems map[string]bool
}

func New() *SortedMap {
	return &SortedMap{
		items:     make(map[string]int),
		order:     make(map[string]int),
		stopItems: make(map[string]bool),
	}
}

func (s *SortedMap) AddItem(item string) {
	if _, isInMap := s.items[item]; isInMap == true {
		s.items[item]++
		return
	}
	s.items[item] = 1
}

func (s *SortedMap) AddStopItem(item string) {
	s.stopItems[item] = true
}

func (s *SortedMap) AddOrder(item string, n int) {
	s.order[item] = n
}

func (s *SortedMap) Top(size int) []string {
	type kv struct {
		key string
		val int
	}

	temp := make([]kv, len(s.items), len(s.items))
	for k, v := range s.items {
		_, isStopWord := s.stopItems[k]
		if !isStopWord {
			temp = append(temp, kv{k, v})
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		a := temp[i].val*10000 + (1000 - s.order[temp[i].key])
		b := temp[j].val*10000 + (1000 - s.order[temp[j].key])

		return a > b
	})

	result := make([]string, size, size)
	for i, kv := range temp {
		if i < size {
			//result = append(result, fmt.Sprintf("%d. %temp: %d [entrance:  %d]\n", i+1, kv.key, kv.val, order[kv.key]))
			if kv.key != "" {
				result[i] = fmt.Sprintf("%s: %d", kv.key, kv.val)
			}
		}
	}

	return result
}

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

func (s *SortedMap) AddItem(item string) int {
	_, ok := s.items[item]
	if ok {
		s.items[item]++
	} else {
		s.items[item] = 1
	}

	return s.items[item]
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

	temp := make([]kv, len(s.items))

	for k, v := range s.items {
		if _, ok := s.stopItems[k]; !ok {
			temp = append(temp, kv{k, v})
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].val > temp[j].val
	})

	temp = temp[0:size]

	sort.Slice(temp, func(i, j int) bool {
		return s.order[temp[j].key] > s.order[temp[i].key]
	})

	result := make([]string, size)

	for i, kv := range temp {
		if i < size {
			if kv.key != "" {
				result[i] = fmt.Sprintf("%s: %d", kv.key, kv.val)
			}
		}
	}

	return result
}

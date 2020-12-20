package sortedmap

import (
	"fmt"
	"sort"
	"sync"
)

type SortedMap struct {
	items     map[string]int
	lineOrder map[string][2]int
	stopItems map[string]bool
	mutex     *sync.Mutex
}

func New() *SortedMap {
	return &SortedMap{
		items:     make(map[string]int),
		lineOrder: make(map[string][2]int),
		stopItems: make(map[string]bool),
		mutex:     new(sync.Mutex),
	}
}

func (s *SortedMap) AddItem(item string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.items[item]
	if ok {
		s.items[item]++
	} else {
		s.items[item] = 1
	}

	return s.items[item]
}

func (s *SortedMap) AddStopItem(item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.stopItems[item] = true
}

func (s *SortedMap) AddLineOrder(item string, lineNumber int, positionInLine int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if o, ok := s.lineOrder[item]; ok {
		if lineNumber < o[0] {
			s.lineOrder[item] = [2]int{lineNumber, positionInLine}
		} else if lineNumber == o[0] && positionInLine < o[1] {
			s.lineOrder[item] = [2]int{lineNumber, positionInLine}
		}
	} else {
		s.lineOrder[item] = [2]int{lineNumber, positionInLine}
	}
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
		if s.lineOrder[temp[j].key][0] > s.lineOrder[temp[i].key][0] {
			return true
		} else if s.lineOrder[temp[j].key][0] == s.lineOrder[temp[i].key][0] {
			if s.lineOrder[temp[j].key][1] > s.lineOrder[temp[i].key][1] {
				return true
			}
		}

		return false
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

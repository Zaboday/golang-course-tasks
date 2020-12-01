package textprocessor

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

type TestSortedMap struct {
	items     map[string]int
	orderLine map[string][2]int
	stopItems map[string]bool
}

func (s *TestSortedMap) AddItem(item string) int {
	if _, isInMap := s.items[item]; isInMap == true {
		s.items[item]++
	}

	s.items[item] = 1

	return s.items[item]
}

func (s *TestSortedMap) AddStopItem(item string) {
	s.stopItems[item] = true
}

func (s *TestSortedMap) AddLineOrder(item string, lineNumber int, linePosition int) {
	s.orderLine[item] = [2]int{lineNumber, linePosition}
}

type MockSortedMap struct {
}

func (s *MockSortedMap) AddItem(item string) int {
	return 1
}

func (s *MockSortedMap) AddStopItem(item string) {
}

func (s *MockSortedMap) AddLineOrder(item string, lineNumber int, linePosition int) {
}

// go test ./pkg/textprocessor
func TestTextProcessor_isStopWords(t *testing.T) {
	cases := []struct {
		prevWord string
		nextWord string
		expected bool
	}{
		{"foo", "bar", false},
		{"foo.", "bar", false},
		{"foo.", "Bar", true},
		{"foo", ".Bar", false},
		{"Foo", ".Bar", false},
		{".", "B", true},
		{"", "Bar", false},
	}

	for i, c := range cases {
		var p TextProcessor
		if p.isStopWords(c.prevWord, c.nextWord) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestTextProcessor_isValidWord(t *testing.T) {
	cases := []struct {
		word     string
		expected bool
	}{
		{",", false},
		{", ", false},
		{",  ", false},
		{",   ", true},
		{"asdasdasda asdasd asdaddas", true},
		{"a", false},
		{"as", false},
		{"asdads", true},
		{"1", false},
		{"1%", false},
		{"11`", false},
		{"11.", false},
		{"111.", true},
	}

	for i, c := range cases {
		var p TextProcessor
		p.wordLength = 3

		if p.isValidWord(c.word) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestTextProcessor_fillStopWordsByLine(t *testing.T) {
	cases := []struct {
		line     string
		expected map[string]bool
	}{
		{"some line", map[string]bool{"some": true, "line": true}},
		{"StaRt Ende", map[string]bool{"start": true, "ende": true}},
		{"StaRt End.", map[string]bool{"start": true}},
		{"foo bar", map[string]bool{}},
		{"start, foo bar was beautiful barend1", map[string]bool{"start,": true, "barend1": true}},
		{" ", map[string]bool{}},
		{"                            ", map[string]bool{}},
		{".,.,.         \n            ../,", map[string]bool{".,.,": true, "../,": true}},
	}

	for i, c := range cases {
		sm := TestSortedMap{map[string]int{}, map[string][2]int{}, map[string]bool{}}
		ss := strings.Fields(c.line)

		var p = New(&sm, 3)

		p.fillStopWordsByLine(ss)

		if !isEqualStopWords(sm.stopItems, c.expected) {
			t.Errorf("Usecase [%d]. expected %v, actual %v", i, c.expected, sm.stopItems)
		}
	}
}

func TestTextProcessor_clearWord(t *testing.T) {
	cases := []struct {
		actual   string
		expected string
	}{
		{"Foo", "foo"},
		{"FOo", "foo"},
		{"FOO.", "foo"},
		{".FOO", ".foo"},
		{".....FOO", ".....foo"},
		{"Foo..", "foo."},
	}

	for i, c := range cases {
		var p TextProcessor
		if p.clearWord(c.actual) != c.expected {
			t.Errorf("Usecase [%d]. clearWord() expected %v, actual %v", i, c.expected, c.actual)
		}
	}
}

// go test -bench=. -benchmem ./pkg/textprocessor

func BenchmarkTextProcessor_ProcessLine(b *testing.B) {
	sm := MockSortedMap{}

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var p = New(&sm, 3)

		file, _ := os.Open("files/src_test.txt")
		sc := bufio.NewScanner(file)

		b.StartTimer()

		for sc.Scan() {
			p.ProcessLine(sc.Text(), 1)
		}
	}
}

// Return true if two maps have same length and same key->values
func isEqualStopWords(a map[string]bool, b map[string]bool) bool {
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

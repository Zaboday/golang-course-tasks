package textprocessor

import (
	"strings"
	"testing"
)

// go test ./pkg/textprocessor
func TestIsStopWords(t *testing.T) {
	useCases := []struct {
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
		{"", "Bar", false}, //6
	}

	for i, c := range useCases {
		var p TextProcessor
		if p.isStopWords(c.prevWord, c.nextWord) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestIsValidWord(t *testing.T) {
	useCases := []struct {
		word     string
		expected bool
	}{
		{",", false},
		{", ", false},
		{",  ", false},
		{",   ", true},
		{"asdasdasda asdasd asdaddas", true},
		{"a", false}, //5
		{"as", false},
		{"asdads", true},
		{"1", false},
		{"1%", false},
		{"11`", false},
		{"11.", false},
		{"111.", true}, //12
	}

	for i, c := range useCases {
		var p TextProcessor
		if p.isValidWord(c.word) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestFillStopWordsByLine(t *testing.T) {
	useCases := []struct {
		line     string
		expected map[string]bool
	}{
		{"some line", map[string]bool{"some": true, "line": true}},
		{"foo bar", map[string]bool{}},
		{"start, foo bar was beautiful barend1", map[string]bool{"start,": true, "barend1": true}},
		{" ", map[string]bool{}},
		{"                            ", map[string]bool{}},
		{".,.,.         \n            ../,", map[string]bool{".,.,.": true, "../,": true}},
	}

	for i, c := range useCases {
		var p TextProcessor
		ss := strings.Fields(c.line)
		p.setStopWords(make(map[string]bool))
		p.fillStopWordsByLine(ss)

		if !isEqualStopWords(p.stopWords, c.expected) {
			t.Errorf("Usecase [%d]. expected %v, actual %v", i, c.expected, p.stopWords)
		}
	}
}

/*func TestGetTop(t *testing.T) {
	words := map[string]int{"f": 1, "b": 2, "h": 2, "d": 5}
	order := map[string]int{"h": 1, "b": 2}
	emptyOrder := make(map[string]bool)

	useCases := []struct {
		words     map[string]int
		order     map[string]int
		stopWords map[string]bool
		topSize   int
		expected  []string
	}{
		{words, make(map[string]int), emptyOrder, 1, []string{"d: 5"}},
		{words, order, emptyOrder, 2, []string{"d: 5", "h: 2"}},
		{words, order, emptyOrder, 3, []string{"d: 5", "h: 2", "b: 2"}},
		{words, make(map[string]int), map[string]bool{"b": true}, 3, []string{"d: 5", "h: 2", "f: 1"}},
		{words, make(map[string]int), map[string]bool{"b": true, "d": true}, 3, []string{"h: 2", "f: 1", ""}},
		{words, map[string]int{"b": 1, "h": 2}, map[string]bool{"d": true}, 5, []string{"b: 2", "h: 2", "f: 1", "", ""}},
	}

	for i, c := range useCases {
		var p TextProcessor
		p.setStopWords(c.stopWords)
		top := p.getTop(c.words, c.order, c.topSize)

		if !isEqualSlices(c.expected, top) {
			t.Errorf("Usecase [%d]. expected %v, actual %v", i, c.expected, top)
		}
	}
}*/

// go test -bench=. -benchmem ./pkg/slicer

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

package textprocessor

import (
	"strings"
	"testing"
)

// go test ./pkg/textprocessor
func TestIsStopWords(t *testing.T) {
	var p Processor
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
		if p.isStopWords(c.prevWord, c.nextWord) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestIsValidWord(t *testing.T) {
	var p Processor
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
		if p.isValidWord(c.word) != c.expected {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

func TestFillStopWordsByLine(t *testing.T) {
	var p Processor
	useCases := []struct {
		line     string
		expected map[string]int
	}{
		{"some line", map[string]int{"some": 1, "line": 1}},
		{"foo bar", map[string]int{}},
		{"start, foo bar was beautiful barend1", map[string]int{"start,": 1, "barend1": 1}},
		{" ", map[string]int{}},
		{"                            ", map[string]int{}},
		{".,.,.         \n            ../,", map[string]int{".,.,.": 1, "../,": 1}},
	}

	for i, c := range useCases {
		sw := make(map[string]int)
		ss := strings.Fields(c.line)
		p.fillStopWordsByLine(ss, sw)
		if !isEqualMaps(sw, c.expected) {
			t.Errorf("Usecase [%d]. expected %v, actual %v", i, c.expected, sw)
		}
	}
}

func TestGetTop(t *testing.T) {
	var p Processor

	words := map[string]int{"f": 1, "b": 2, "h": 2, "d": 5}
	emptyMap := make(map[string]int)

	useCases := []struct {
		words     map[string]int
		order     map[string]int
		stopWords map[string]int
		topSize   int
		expected  []string
	}{
		{words, emptyMap, emptyMap, 1, []string{"1. d: 5"}},
		{words, emptyMap, emptyMap, 2, []string{"1. d: 5", "2. h: 2"}},
		{words, emptyMap, emptyMap, 3, []string{"1. d: 5", "2. h: 2", "3. b: 2"}},
	}
	for i, c := range useCases {
		top := p.getTop(c.words, c.order, c.stopWords, c.topSize)
		if !isEqualSlices(top, c.expected) {
			t.Errorf("Usecase [%d]. expected %v, actual %v", i, c.expected, top)
		}
	}

}

// Return true if two maps have same length and same key->values
func isEqualMaps(a map[string]int, b map[string]int) bool {
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

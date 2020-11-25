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
		//{"foo bar", map[string]int{}},
	}

	for i, c := range useCases {
		sw := make(map[string]int)
		ss := strings.Fields(c.line)
		p.fillStopWordsByLine(ss, sw)
		if !isEqualMaps(sw, c.expected) {
			t.Errorf("Usecase [%d]. expected %v", i, c.expected)
		}
	}
}

// Return true if two maps have same length and same key, values
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

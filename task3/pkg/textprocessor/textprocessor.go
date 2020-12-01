package textprocessor

import (
	"regexp"
	"strings"
	"sync"
	"unicode"
)

type TextProcessor struct {
	sortedMap  SortedWordsMap
	wordLength int
	regexp     *regexp.Regexp
	mu         sync.Mutex
}

const POINT = 46

type SortedWordsMap interface {
	AddItem(item string) int
	AddStopItem(item string)
	AddLineOrder(item string, l int, n int)
}

func New(sm SortedWordsMap, l int) *TextProcessor {
	r := regexp.MustCompile(`[^a-zA-Z0-9\.\s]`)

	return &TextProcessor{
		sortedMap:  sm,
		wordLength: l,
		regexp:     r,
	}
}

func (p *TextProcessor) ProcessLine(line string, lineNumber int) {
	if line == "" {
		return
	}

	line = p.regexp.ReplaceAllString(line, "")
	words := strings.Fields(line)
	p.fillStopWordsByLine(words)

	prevWordOrigin := ""

	i := 0
	for _, word := range words {
		if p.isStopWords(prevWordOrigin, word) {
			p.sortedMap.AddStopItem(p.clearWord(word))
			p.sortedMap.AddStopItem(p.clearWord(prevWordOrigin))
		} else {
			prevWordOrigin = word
			word = p.clearWord(word)
			if p.isValidWord(word) {
				p.sortedMap.AddItem(word)
				p.sortedMap.AddLineOrder(word, lineNumber, i)
			}
		}
		i++
	}
}

// Determines two words are begin and end of sentence.
func (p *TextProcessor) isStopWords(prevWord string, nextWord string) bool {
	if prevWord == "" || nextWord == "" {
		return false
	}

	if prevWord[len(prevWord)-1] == POINT {
		r := []rune(nextWord)
		if unicode.IsUpper(r[0]) {
			return true
		}
	}

	return false
}

// Determines valid word.
func (p *TextProcessor) isValidWord(word string) bool {
	return len(word) > p.wordLength
}

// Determines stop-words in line. First and last words of line are stop-words.
func (p *TextProcessor) fillStopWordsByLine(line []string) {
	if len(line) == 0 {
		return
	}

	firstWord := p.clearWord(line[0])

	if p.isValidWord(firstWord) {
		p.sortedMap.AddStopItem(strings.ToLower(firstWord))
	}

	lastWord := p.clearWord(line[len(line)-1])
	if p.isValidWord(lastWord) {
		p.sortedMap.AddStopItem(strings.ToLower(lastWord))
	}
}

// Removes "." from the end of word
func (p *TextProcessor) clearWord(word string) string {
	if word == "" {
		return word
	}

	word = strings.ToLower(word)

	if word[len(word)-1] == POINT {
		r := []rune(word)
		return string(r[0 : len(r)-1])
	}

	return word
}

package textprocessor

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type TextProcessor struct {
	sortedMap  SortedWordsMap
	counter    int
	wordLength int
	regexp     *regexp.Regexp
}

type SortedWordsMap interface {
	AddItem(item string)
	AddStopItem(item string)
	AddOrder(item string, n int)
}

func New(sm SortedWordsMap) *TextProcessor {
	r, err := regexp.Compile(`[^a-zA-Z0-9\.\s]`)
	if err != nil {
		panic(fmt.Sprintf("Regexp compilation error: %s", err))
	}

	return &TextProcessor{
		sortedMap:  sm,
		wordLength: 3,
		regexp:     r,
	}
}

func (p *TextProcessor) SetWordLength(len int) {
	p.wordLength = len
}

func (p *TextProcessor) ProcessLine(line string) {
	if line == "" {
		return
	}

	line = p.regexp.ReplaceAllString(line, "")
	words := strings.Fields(line)
	p.fillStopWordsByLine(words)

	prevWordOrigin := ""
	for _, word := range words {

		if p.isStopWords(prevWordOrigin, word) {
			p.sortedMap.AddStopItem(p.clearWord(word))
			p.sortedMap.AddStopItem(p.clearWord(prevWordOrigin))
			continue
		}

		prevWordOrigin = word
		word = p.clearWord(word)
		if p.isValidWord(word) {
			p.sortedMap.AddItem(word)
			p.sortedMap.AddOrder(word, p.counter)
		}

		p.counter++
	}
}

// Determines two words are begin and end of sentence.
func (p *TextProcessor) isStopWords(prevWord string, nextWord string) bool {
	if prevWord == "" || nextWord == "" {
		return false
	}
	// 46 is "."
	if prevWord[len(prevWord)-1] == 46 {
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
	// 46 is "."
	if word[len(word)-1] == 46 {
		r := []rune(word)
		return string(r[0 : len(r)-1])
	}
	return word
}

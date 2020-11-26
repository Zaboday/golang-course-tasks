package textprocessor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

type TextProcessor struct {
	file      io.Reader
	stopWords map[string]bool
}

func New(reader io.Reader) *TextProcessor {
	return &TextProcessor{
		file:      reader,
		stopWords: make(map[string]bool),
	}
}

func (p *TextProcessor) Process(fileNameTxt string) {
	result := make(map[string]int)
	order := make(map[string]int)

	lines, err := p.getLinesFromFile(fileNameTxt)
	if err != nil {
		panic(err)
	}

	i := 0
	for _, line := range lines {

		lineWords := strings.Fields(line)
		p.fillStopWordsByLine(lineWords)

		prevWord := ""
		for _, word := range lineWords {

			if prevWord != "" && p.isStopWords(prevWord, word) {
				p.stopWords[word] = true
				p.stopWords[prevWord] = true
			}

			_, isInResult := result[word]
			if isInResult {
				result[word]++
				continue
			}

			if p.isValidWord(word) {
				order[word] = i
				result[word] = 1
			}

			prevWord = word
			i++
		}
	}

	for _, v := range p.getTop(result, order, 10) {
		fmt.Println(v)
	}
}

// Determines whether two words are end and begin of sentence.
func (p *TextProcessor) isStopWords(prevWord string, currWord string) bool {
	if prevWord == "" || currWord == "" {
		return false
	}

	if prevWord[len(prevWord)-1] == 46 {
		r := []rune(currWord)
		if unicode.IsUpper(r[0]) {
			return true
		}
	}

	return false
}

func (p *TextProcessor) setStopWords(stopWords map[string]bool) {
	p.stopWords = stopWords
}

// Determines stopWords in line and adds to map.
func (p *TextProcessor) fillStopWordsByLine(line []string) {
	if len(line) == 0 {
		return
	}

	if p.isValidWord(line[0]) {
		p.stopWords[line[0]] = true
	}

	if p.isValidWord(line[len(line)-1]) {
		p.stopWords[line[len(line)-1]] = true
	}
}

func (p *TextProcessor) isValidWord(word string) bool {
	return len(word) > 3
}

func (p *TextProcessor) getLinesFromFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, nil
}

func (p *TextProcessor) getTop(words map[string]int, order map[string]int, topSize int) []string {
	type kv struct {
		Key   string
		Value int
	}

	s := make([]kv, len(words), len(words))
	for k, v := range words {
		_, isStopWord := p.stopWords[k]
		if !isStopWord {
			s = append(s, kv{k, v})
		}
	}

	sort.Slice(s, func(i, j int) bool {
		// Сортируем в соответствии с позицией первого вхождения.
		a := s[i].Value*10000 + (1000 - order[s[i].Key])
		b := s[j].Value*10000 + (1000 - order[s[j].Key])

		return a > b
	})

	result := make([]string, topSize, topSize)
	for i, kv := range s {
		if i < topSize {
			//result = append(result, fmt.Sprintf("%d. %s: %d [entrance:  %d]\n", i+1, kv.Key, kv.Value, order[kv.Key]))
			if kv.Key != "" {
				result[i] = fmt.Sprintf("%s: %d", kv.Key, kv.Value)
			}
		}
	}

	return result
}

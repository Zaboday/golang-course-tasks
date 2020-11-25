package textprocessor

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type Processor struct {
}

func (p *Processor) ProcessTxt(fileNameTxt string) {
	result := make(map[string]int)
	stopWords := make(map[string]int)
	order := make(map[string]int)

	lines, err := p.getLinesFromFile(fileNameTxt)
	if err != nil {
		panic(err)
	}

	i := 0
	for _, line := range lines {

		lineWords := strings.Fields(line)
		p.fillStopWordsByLine(lineWords, stopWords)

		prevWord := ""
		for _, word := range lineWords {

			if prevWord != "" && p.isStopWords(prevWord, word) {
				stopWords[word] = 1
				stopWords[prevWord] = 1
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
	p.showTop(result, order, stopWords, 10)
}

// Determines whether two words are end and begin of sentence.
func (p *Processor) isStopWords(prevWord string, currWord string) bool {
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

// Determines stopWords in line and adds to map.
func (p *Processor) fillStopWordsByLine(line []string, stopWords map[string]int) {
	if len(line) == 0 {
		return
	}

	if p.isValidWord(line[0]) {
		stopWords[line[0]] = 1
	}

	if p.isValidWord(line[len(line)-1]) {
		stopWords[line[len(line)-1]] = 1
	}
}

func (p *Processor) isValidWord(word string) bool {
	return len(word) > 3
}

func (p *Processor) getLinesFromFile(path string) ([]string, error) {

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

func (p *Processor) showTop(words map[string]int, order map[string]int, stopWords map[string]int, topSize int) {
	type kv struct {
		Key   string
		Value int
	}

	s := make([]kv, len(words), len(words))
	for k, v := range words {
		_, isStopWord := stopWords[k]
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

	for i, kv := range s {
		if i < topSize {
			fmt.Printf("%d. %s: %d [entrance:  %d]\n", i+1, kv.Key, kv.Value, order[kv.Key])
		}
	}
}

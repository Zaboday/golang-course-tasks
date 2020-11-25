package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fileName := "./files/src.txt"
	result := make(map[string]int)
	stopWords := make(map[string]int)

	lines, err := getLinesFromFile(fileName)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {

		lineWords := strings.Fields(line)
		fillStopWordsByLine(lineWords, stopWords)

		for _, word := range lineWords {
			//var prevWord string
			_, isInResult := result[word]

			if isInResult {
				result[word]++
				continue
			}

			if isValidWord(word) {
				result[word] = 1
			}
		}
	}

	showTop(result, stopWords, 10)
}

func fillStopWordsByLine(line []string, stopWords map[string]int) {
	if len(line) == 0 {
		return
	}

	if isValidWord(line[0]) {
		stopWords[line[0]] = 1
	}

	if isValidWord(line[len(line)-1]) {
		stopWords[line[len(line)-1]] = 1
	}
}

func isValidWord(word string) bool {
	return len(word) > 3
}

func getLinesFromFile(path string) ([]string, error) {

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

func showTop(words map[string]int, stopWords map[string]int, topSize int) {
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
		return s[i].Value > s[j].Value
	})

	for i, kv := range s {
		if i < topSize {
			fmt.Printf("%d. %s: %d\n", i+1, kv.Key, kv.Value)
		}
	}
}

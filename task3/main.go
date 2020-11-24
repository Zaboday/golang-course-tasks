package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//result := make(map[string]int)
	fileName := "./files/src.txt"
	/*words, err := getWord(fileName)
	if err != nil {
		panic(err)
	}
	for _, word := range words {

		if !isValidWord(word) {
			continue
		}
		count, isOk := result[word]

		if isOk != true {
			result[word] = count
			continue
		}

		result[word]++
	}*/
	lines, err := getLines(fileName)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func isValidWord(w string) bool {
	return len(w) > 3
}

func getLines(path string) ([]string, error) {

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

func getWord(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)

	var words []string

	for sc.Scan() {
		words = append(words, sc.Text())
	}

	return words, nil
}

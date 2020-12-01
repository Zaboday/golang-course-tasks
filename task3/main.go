package main

import (
	"bufio"
	"flag"
	"fmt"
	"main/pkg/sortedmap"
	"main/pkg/textprocessor"
	"os"
	"sync"
)

func main() {
	var fileName string

	var wordLength int

	var topSize int

	flag.StringVar(&fileName, "f", "files/src.txt", "file path with text")
	flag.IntVar(&wordLength, "wl", 3, "word length")
	flag.IntVar(&topSize, "ts", 10, "size of top list words")
	flag.Parse()

	var sm = sortedmap.New()
	p := textprocessor.New(sm, wordLength)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File open error: " + err.Error())
		return
	}
	defer file.Close()

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	sc := bufio.NewScanner(file)
	i := 0
	for sc.Scan() {
		wg.Add(1)
		i++
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			p.ProcessLine(sc.Text(), i)
			mu.Unlock()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

	for _, v := range sm.Top(topSize) {
		fmt.Println(v)
	}
}

package main

import (
	"main/pkg/textprocessor"
)

func main() {
	fileName := "./files/src_test1.txt"
	var processor textprocessor.Processor

	processor.ProcessTxt(fileName)
}

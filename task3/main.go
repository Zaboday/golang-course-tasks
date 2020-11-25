package main

import (
	"main/pkg/textprocessor"
)

func main() {
	fileName := "./files/src.txt"
	var processor textprocessor.Processor

	processor.ProcessTxt(fileName)
}

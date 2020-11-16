package main

import (
	"fmt"
	"main/pkg/slicer"
)

func main() {
	var sorted []int
	var x int

	for {
		fmt.Scan(&x)

		if x == 999 {
			// Exit app
			fmt.Println(sorted)
			return
		}

		if x >= 0 {
			sorted = slicer.Insert(x, sorted)
		} else {
			sorted = slicer.Delete(-x, sorted)
		}

		fmt.Println(sorted)
	}

}

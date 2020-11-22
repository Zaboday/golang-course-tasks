package main

import (
	"fmt"
	"main/pkg/linkedlist"
	//"main/pkg/slicer"
)

func main() {
	callLinkedList()
}

func callLinkedList() {
	var l linkedlist.List
	var x int

	for {
		fmt.Scan(&x)

		if x == 999 {
			// Exit app
			fmt.Println(l.DisplayChain())
			return
		}

		if x >= 0 {
			l.Insert(x)
		} else {
			l.Delete(-x)
		}

		fmt.Println(l.DisplayChain())
	}

}

/*func callSliceManager() {
	var sorted []int
	var manager slicer.SliceManager
	var x int

	for {
		fmt.Scan(&x)

		if x == 999 {
			// Exit app
			fmt.Println(sorted)
			return
		}

		if x >= 0 {
			sorted = manager.Insert(x, sorted)
		} else {
			sorted = manager.Delete(-x, sorted)
		}

		fmt.Println(sorted)
	}
}*/

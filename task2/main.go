package main

import (
	"fmt"
	//"main/pkg/linkedlist"
	"main/pkg/slicer"
)

func main() {
	callSliceManager()
}

/*func callLinkedList() {
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

}*/

func callSliceManager() {
	var result []int
	var manager slicer.SliceManager
	var x int

	for {
		fmt.Scan(&x)

		if x == 999 {
			// Exit app
			fmt.Println("Exit code")
			return
		}

		if x >= 0 {
			result = manager.Insert(x)
		} else {
			result = manager.Delete(-x)
		}

		fmt.Println(result)
	}
}

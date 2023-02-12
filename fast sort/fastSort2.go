package main

import "fmt"

func selectionSortFast(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func main() {
	loga := []int{5, 3, 6, 2, 10}
	selectionSortFast(loga)
	fmt.Println(loga)
}

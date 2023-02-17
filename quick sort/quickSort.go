package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		less := make([]int, 0)
		more := make([]int, 0)
		for _, v := range arr[1:] {
			if v < pivot {
				less = append(less, v)
			} else if v > pivot {
				more = append(more, v)
			}
		}
		return append(append(quickSort(less), pivot), quickSort(more)...)

	}
}

func main() {
	mas := []int{1, 78, 2, 64, 85, 35, 7, 900}
	fmt.Println(quickSort(mas))
}

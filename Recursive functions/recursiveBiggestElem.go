package main

import "fmt"

func recursiveBiggest(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	} else {
		subMax := recursiveBiggest(arr[1:])
		if arr[0] > subMax {
			return arr[0]
		} else {
			return subMax
		}
	}
}

func main() {
	arr := []int{1, 9, 2, 5, 3, 6, 88}
	fmt.Println(recursiveBiggest(arr))
}

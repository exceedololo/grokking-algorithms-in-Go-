package main

import "fmt"

func recursiveAmmount(arr []int) int {
	l := len(arr)
	if l == 0 {
		return 0
	} else {
		return 1 + recursiveAmmount(arr[1:])
	}
}

func main() {
	arr := []int{1, 4, 2, 5}
	fmt.Println(recursiveAmmount(arr))
}

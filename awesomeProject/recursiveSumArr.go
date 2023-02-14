package main

import "fmt"

func summArr(arr []int) int {
	lenght := len(arr)
	if lenght == 0 {
		return 0 //arr[lenght]
	} else {
		return arr[0] + summArr(arr[1:lenght])
	}
}

func main() {
	arra := []int{1, 5, 4, 6}
	fmt.Println(summArr(arra))

}

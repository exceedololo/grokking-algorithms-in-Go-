package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int
	for i := range arr {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
		i--
	}
	return newArr

}

func main() {
	loga := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(loga))
}

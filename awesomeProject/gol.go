package main

import (
	"errors"
	"fmt"
	"sort"
)

func binarySearch(list []int, item int) (int, error) {
	low := 0
	high := len(list) - 1

	sort.Ints(list)
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return 0, errors.New("not found")
}

func main() {
	fmt.Println("ti kurica")
	puk := []int{1, 4, 2, 5, 6, 4, 2, 7, 9, 9, 3}
	index, err := binarySearch(puk, 2)
	if err != nil {
		fmt.Println("item not found")
	} else {
		fmt.Printf("item is at number %d\n", index)
	}

}

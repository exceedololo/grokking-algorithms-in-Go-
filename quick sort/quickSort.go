package main

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
		return (quickSort(less) + pivot) + quickSort(more)

	}
}

func main() {

}

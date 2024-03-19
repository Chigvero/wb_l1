package main

import (
	"fmt"
)

func main() {

	array := []int{1001, 1000, 1001, 12, 14, 1, 1002, 100, 321, 21, 1001, 1002, 1000, 50, 999, 998, 1, 13}
	fmt.Println(array)
	quickSort(&array, 0, len(array)-1)
	//swap(&array)
	fmt.Println(array)
}

func quickSort(array *[]int, low, high int) {

	if low < high {
		pivot := partition(array, low, high)
		quickSort(array, low, pivot-1)
		quickSort(array, pivot+1, high)

	}

}

func partition(array *[]int, low, high int) int {
	pivot := (*array)[high]
	i := low
	for j := low; j < high; j++ {

		if (*array)[j] < pivot {
			(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
			i++
		}
		//fmt.Println(j, i, ":", array)
	}

	(*array)[i], (*array)[high] = (*array)[high], (*array)[i]
	return i
}

// func quickSort()

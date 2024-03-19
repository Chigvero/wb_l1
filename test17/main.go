package main

import "fmt"

func binary_finder(arr []int, value int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (left + right) / 2
		switch {
		case arr[middle] < value:
			left = middle + 1
		case arr[middle] > value:
			right = middle - 1
		case arr[middle] == value:
			return middle
		}
	}
	return -1
}

func main() {
	array := []int{0, 6, 11, 11, 18, 25, 28, 28, 37, 40, 45, 47, 47, 56, 58, 59, 62, 66, 74, 81, 81, 87, 89, 94, 95}
	fmt.Println(binary_finder(array, 28))
}

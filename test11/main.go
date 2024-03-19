package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	fmt.Println(intersect(slice1, slice2))
}

func intersect(slice1 []int, slice2 []int) []int {
	var slice []int
	for _, i := range slice1 {
		for _, i2 := range slice2 {
			if i == i2 {
				slice = append(slice, i)
			}
		}
	}
	return slice
}

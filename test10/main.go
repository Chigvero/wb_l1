package main

import (
	"fmt"
)

func main() {
	tempMap := make(map[float32][]float32)
	var x int
	tempArray := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//fmt.Print(tempArray)
	for _, v := range tempArray {
		x = int(v / 10)
		z := float32(x * 10)
		tempMap[z] = append(tempMap[z], v)
	}
	fmt.Print(tempMap)
}

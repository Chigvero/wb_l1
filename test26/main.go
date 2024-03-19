package main

import (
	"fmt"
)

func main() {
	fmt.Print("START")
	var s string
	fmt.Scan(&s)
	myR := []byte(s)
	myMap := make(map[byte]int)
	for _, v := range myR {
		if v > 64 && v < 91 {
			myMap[v+32]++
		} else {
			myMap[v]++
		}
	}
	i := true
	for _, v := range myMap {
		if v > 1 {
			i = false
		}
	}
	fmt.Print(i)
}

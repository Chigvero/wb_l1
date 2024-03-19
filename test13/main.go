package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")

	x := 10
	y := 11
	x = x + y
	y = x - y
	x = x - y
	x, y = y, x
	fmt.Print(x, y)
}

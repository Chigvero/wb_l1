package main

import (
	"fmt"
)

var s string = "ГЛАВРЫБ1"

func main() {
	rs := []rune(s)
	fmt.Println(rs)
	z := len(rs) - 1
	output := []rune{}
	for i, _ := range rs {
		output = append(output, rs[z-i])
	}
	fmt.Print("\n")
	fmt.Println(string(output))
}

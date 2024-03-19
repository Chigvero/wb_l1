package main

import (
	"fmt"
)

func main() {
	arrayS := []string{"cat", "cat", "dog", "cat", "tree"}
	mapS := make(map[string]int)
	for _, v := range arrayS {
		mapS[v] += 1
	}
	fmt.Print(mapS)
}

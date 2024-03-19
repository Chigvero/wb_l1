package main

import (
	"fmt"
)

var justString string

func createHugeString(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("Invalid length")
	}

	chars := make([]byte, length)
	for i := range chars {
		chars[i] = 'a'
	}
	return string(chars), nil
}

func someFunc() {
	v, err := createHugeString(1 << 10)
	if err != nil {
		fmt.Println("Error creating huge string:", err)
		return
	}

	if len(v) < 100 {
		fmt.Println("Error: length of huge string is less than 100")
		return
	}

	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println("justString:", justString)
}

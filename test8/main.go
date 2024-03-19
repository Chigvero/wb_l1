package main

import (
	"fmt"
	"strconv"
)

func IntToBits(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func main() {
	var i, v, ch int
	fmt.Println("enter number")
	fmt.Scan(&ch)
	s := IntToBits(ch)
	fmt.Print(s)
	slice := []byte(s)
	fmt.Println("enter i and s[i]")
	fmt.Scan(&i)
	fmt.Scan(&v)
	if v == 1 {
		slice[i] = 49
	} else {
		slice[i] = 48
	}
	newVal, err := strconv.ParseInt(string(slice), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal)
}

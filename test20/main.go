package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("START")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	s = s[0 : len(s)-1]
	if err != nil {
		panic(err)
	}

	myStrings := strings.Split(s, " ")
	myStrings1 := []string{}
	for i, _ := range myStrings {
		myStrings1 = append(myStrings1, myStrings[len(myStrings)-1-i])
	}
	for _, v := range myStrings1 {
		fmt.Print(v, " ")
	}
}

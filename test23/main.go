package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	z, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	z = z[0 : len(z)-1]
	i, _ := strconv.Atoi(z)
	mystring := "AKSKA AKA "
	m1 := []rune(mystring[0:i])
	m2 := []rune(mystring[i+1:])
	for j := 0; j < len(m2); j++ {
		m1 = append(m1, m2[j])
	}
	//m1=append(m1, m2...)
	fmt.Print(string(m1))
}

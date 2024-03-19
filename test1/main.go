package main

import (
	"fmt"
)

type Human struct {
	Name       string
	SecondNAME string
	Age        int
}

func (H *Human) Hi() {
	fmt.Println("hi")

}
func (H *Human) Intro() {
	fmt.Printf("My Name is %v %v\n", H.Name, H.SecondNAME)
}

type Action struct {
	Human
}

func (A *Action) walk() {
	fmt.Println("walk,walk,walk")
}

func main() {
	fmt.Println("START")
	a := Action{}
	a.Hi()
	a.walk()

}

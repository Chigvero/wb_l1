package main

import (
	"fmt"
	"time"
)

func main() {
	exit := false
	go func() {
		for {
			switch {
			case exit:
				fmt.Println("Goroutine termination signal received")
				return
			default:
				fmt.Println("Goroutine works...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	exit = true
	time.Sleep(time.Second)
	fmt.Println("Program completed")
}

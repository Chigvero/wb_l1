package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan bool)
	go func() {
		for {
			select {
			case <-exit:
				fmt.Println("Goroutine termination signal received")
				return
			default:
				fmt.Println("Goroutine works...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)
	exit <- true
	close(exit)
	fmt.Println("Program completed")
}

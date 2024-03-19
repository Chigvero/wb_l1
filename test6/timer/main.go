package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Goroutine termination signal received")
				runtime.Goexit()
			default:
				fmt.Println("Goroutine works...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 6)
	fmt.Println("Program completed")
}

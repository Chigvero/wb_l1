package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main() {
// 	fmt.Print("START")
// 	var (
// 		N  time.Duration
// 		wg sync.WaitGroup
// 	)
// 	fmt.Scan(&N)

// 	channel := make(chan string)
// 	go func() {
// 		for {
// 			wg.Add(1)
// 			go input(&wg, channel)

// 			go fmt.Println(<-output(channel))
// 		}
// 	}()
// 	time.Sleep(N * time.Second)
// }

func input(wg *sync.WaitGroup, channel chan string) chan string {

	defer wg.Done()
	x := ""
	fmt.Scan(&x)
	channel <- x
	return channel

}

func output(channel chan string) <-chan string {
	return channel
}

func Sender(mychan chan int) {
	for {
		mychan <- rand.Intn(1000)
	}
}

func Receiver(mychan chan int) {
	for v := range mychan {
		fmt.Println(v)
	}
}

func main() {
	var n time.Duration
	fmt.Scan(&n)

	mychan := make(chan int)

	go Receiver(mychan)
	go Sender(mychan)

	time.Sleep(n * time.Second)
	close(mychan)
}

package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	count int
}

func (C *Counters) increment() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		C.count++
	}()
	wg.Wait()
}

func main() {
	fmt.Print("START")
	Counter := Counters{}
	for i := 0; i < 1000000; i++ {
		Counter.increment()
	}
	fmt.Println(Counter.count)
}

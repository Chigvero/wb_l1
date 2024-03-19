package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("START")
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	array := []int{2, 4, 6, 8, 10}
	sum := 0
	for _, i := range array {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			sum += (i * i)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Print(sum)
}

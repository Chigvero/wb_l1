package main

import (
	"fmt"
	"sync"
)

const N = 100

func main() {
	fmt.Print("Start")
	myMap := NewMyMutex()
	mut := sync.Mutex{}
	fmt.Print("Please insert gorutines count :")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < N; i++ {
			wg.Add(1)
			go func(in int) {
				mut.Lock()
				myMap.m[in] = in * in
				mut.Unlock()
				wg.Done()
			}(i)

		}
	}()
	wg.Wait()
	fmt.Print(myMap)
}

type MyMutex struct {
	// sync.Mutex
	m map[int]int
}

func NewMyMutex() *MyMutex {
	return &MyMutex{
		m: make(map[int]int),
	}
}

// type Counters struct {
// 	sync.Mutex
// 	m     map[int]int
// 	count int
// }

// func NewCounters(N int) Counters {
// 	mapa := Counters{
// 		m:     make(map[int]int),
// 		count: N,
// 	}

// 	mapa.m[1] = 0

// 	return mapa
// }

// func main() {
// 	mapa := NewCounters(123)
// 	// mapa.m[1] = 0
// 	fmt.Print(mapa.m)
// }

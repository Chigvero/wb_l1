package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		N int
	)
	Nchan := make(chan int, 100)
	sigChan := make(chan os.Signal, 1)

	fmt.Println("START")

	fmt.Print("insert N:")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		go func(number int) {
			for v := range Nchan {
				fmt.Printf("%d) %d\n", number, v)
			}
		}(i)
	}

	for i := 0; ; i++ {
		select {
		case <-sigChan:
			close(Nchan)
		default:
			if i%10 == 0 {
				time.Sleep(time.Second * 1)
			}
			Nchan <- rand.Intn(1000 - 1 + 1)
		}

	}

	// for {
	// 	fmt.Scan(&insertdata)
	// 	Nchan <- insertdata
	// }

}

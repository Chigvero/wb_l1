package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	inputC := make(chan int)
	outputC := make(chan int, 1)
	go func() {
		for {
			inputC <- rand.Intn(100)
		}
	}()
	go func() {
		for v := range inputC {
			outputC <- (v * 2)
		}
	}()
	for v := range outputC {
		fmt.Println(v)
	}

	select {
	case <-sigChan:
		close(inputC)
		close(outputC)
		os.Exit(1)
	default:
		<-sigChan
	}
}

// func main() {
// 	// Создаем каналы
// 	numbers := make(chan int)
// 	results := make(chan int)
// 	// Горутина для записи чисел в первый канал
// 	go func() {
// 		// Пример массива чисел
// 		for i := 0; ; i++ {
// 			arr := rand.Intn(100)
// 			numbers <- arr
// 			if i%10 == 0 {
// 				time.Sleep(time.Second * 1)
// 			}
// 		}

// 	}()
// 	// Горутина для умножения чисел из первого канала и записи результата во второй канал
// 	go func() {
// 		for num := range numbers {
// 			results <- num * 2
// 		}
// 	}()
// 	// Вывод результатов из второго канала в stdout
// 	for result := range results {
// 		fmt.Println(result)
// 	}
// }

package main

import (
	"time"
)

func my_sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	my_sleep(time.Second * 1)

}

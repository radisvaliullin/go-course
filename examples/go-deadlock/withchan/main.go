package main

import (
	"time"
)

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	outCh := make(chan int)

	go func() {
		for {
			<-ch1
			ch2 <- 0
			outCh <- 0
		}
	}()

	go func() {
		for {
			<-ch2
			ch1 <- 0
			outCh <- 0
		}
	}()

	for {
		select {
		case <-outCh:
			return
		default:
			time.Sleep(time.Second * 2)
		}
	}
}

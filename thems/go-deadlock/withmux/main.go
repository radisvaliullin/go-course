package main

import (
	"fmt"
	"sync"
	"time"
)

var aMux sync.Mutex
var a int

var bMux sync.Mutex
var b int

func main() {

	ch := make(chan int, 1000)

	go fn1(ch)
	go fn2(ch)

	tk5 := time.NewTicker(time.Second * 5)
	for {
		select {
		case i := <-ch:
			fmt.Println("a - ", i)
		case <-tk5.C:
			fmt.Println("tick 5 sec")
		}
	}
}

func fn1(ch chan<- int) {
	for {
		// a
		aMux.Lock()
		a++
		ch <- a
		aMux.Unlock()

		time.Sleep(time.Millisecond * 500)

		// b
		bMux.Lock()
		b++
		bMux.Unlock()

		time.Sleep(time.Millisecond * 500)

		// a and b
		aMux.Lock()
		bMux.Lock()
		a++
		b++
		bMux.Unlock()
		aMux.Unlock()
	}
}

func fn2(ch chan<- int) {
	for {
		// a
		aMux.Lock()
		a++
		aMux.Unlock()

		time.Sleep(time.Millisecond * 500)

		// b
		bMux.Lock()
		b++
		bMux.Unlock()

		time.Sleep(time.Millisecond * 500)

		// a and b
		aMux.Lock()
		bMux.Lock()
		a++
		b++
		aMux.Unlock()
		bMux.Unlock()

	}
}

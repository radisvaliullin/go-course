package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	// Simple Example how run and stop goroutine in true way.
	// We use different kind approaches: send signals or close channels.

	// chans for send messages (data) between goroutine
	inCh := make(chan int, 100)
	toPrint := make(chan int, 100)

	// Run goroutines, first need start counsumers and last producers.

	// Run printer (simulate consumer,
	// for simplicity we only print data in real we can write to db or send to network)
	ptDoneSig := make(chan struct{}, 1)
	go printer(toPrint, ptDoneSig)

	// do some calculate (data modification),
	// read from producer chan and result send to printer (consumer)

	// parallel thread limit
	mp := runtime.GOMAXPROCS(0)

	wg := sync.WaitGroup{}
	for i := 0; i < mp; i++ {
		wg.Add(1)
		go modifyValue(inCh, toPrint, &wg)
	}

	// generate data (producer)
	// in our simple example we just send 1 to chan,
	// in read we cant get this data for example from network
	genStopSig := make(chan struct{}, 1)
	genDoneSig := make(chan struct{}, 1)

	// don't use anonymous function in your code, i used it only for an example
	go func(stopSig <-chan struct{}, doneSig chan<- struct{}) {
		defer func() { doneSig <- struct{}{} }()
		for {
			select {
			case <-stopSig:
				return
			case inCh <- 1:
			default:
				time.Sleep(time.Microsecond * 100)
			}
		}
	}(genStopSig, genDoneSig)

	// After some time wrong something, we must stop our app
	// First we must stop producer and only in final stop consumer
	// Before stoping we must guarantee that all data in channel was handled
	time.Sleep(time.Second * 5)

	// first stop producer
	genStopSig <- struct{}{}
	<-genDoneSig

	// after we can close channel for stoping data modify goroutine
	close(inCh)
	wg.Wait()

	// after we can close channel for stoping consumer (printer)
	close(toPrint)
	<-ptDoneSig

	// after geting last stop signal we can out
	fmt.Println("app done.")
}

//
func modifyValue(in <-chan int, toPrint chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	cntr := 0

	for i := range in {
		cntr += i
		time.Sleep(time.Microsecond * 500)
		toPrint <- cntr
	}
}

//
func printer(toPrint <-chan int, done chan<- struct{}) {

	for i := range toPrint {
		fmt.Printf("i - %v printed\n", i)
	}

	done <- struct{}{}
}

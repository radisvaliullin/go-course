package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	inCh := make(chan int, 100)
	toPrint := make(chan int, 100)

	// generate data
	genStopSig := make(chan struct{}, 1)
	genDoneSig := make(chan struct{}, 1)
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

	// run printer
	ptDoneSig := make(chan struct{}, 1)
	go printer(toPrint, ptDoneSig)

	// do calculate
	// parallel thread limit
	mp := runtime.GOMAXPROCS(0)

	wg := sync.WaitGroup{}
	for i := 0; i < mp; i++ {
		wg.Add(1)
		go modifyValue(inCh, toPrint, &wg)
	}

	// after some time wrong something, we must stop our app
	time.Sleep(time.Second * 5)

	// stop process
	genStopSig <- struct{}{}
	<-genDoneSig
	close(inCh)

	wg.Wait()
	close(toPrint)

	<-ptDoneSig

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

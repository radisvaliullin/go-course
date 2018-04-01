package main

import (
	"fmt"
	"runtime"
	"sync"
)

type store struct {
	mx    sync.Mutex
	store map[string]int
}

func newStore() *store {
	return &store{
		store: map[string]int{},
	}
}

func (s *store) isExist(k string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	_, ok := s.store[k]
	return ok
}

func (s *store) add(k string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.store[k]++
}

func (s *store) findDoubleAdd() bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	for k, v := range s.store {
		if v > 1 {
			fmt.Println("double: ", k, v)
			return true
		}
	}
	return false
}

func main() {

	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	s := newStore()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setStore(s, wg)
	}

	wg.Wait()
	d := s.findDoubleAdd()
	fmt.Println("double: ", d)
}

func setStore(s *store, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		k := fmt.Sprintf("%v", i)
		if !s.isExist(k) {
			s.add(k)
		}
	}
}

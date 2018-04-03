package main

import "fmt"

type getStrInt interface {
	getStr() string
	getInt() int
}

type some struct {
	s string
	i int
}

func (s some) getStr() string {
	return "Some Str"
}

func (s some) getInt() int {
	return 73
}

func main() {

	var i getStrInt

	s := some{}

	i = &s

	fmt.Print(i)
}

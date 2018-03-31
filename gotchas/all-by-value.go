package main

import "fmt"

func main() {

	i := 73
	pi := &i

	fmt.Printf("i - %v, *i - %v\n", i, &i)
	fmt.Printf("pi - %v, *pi - %v\n", pi, &pi)

	takePointer(pi)
}

func takePointer(pi *int) {
	fmt.Printf("takePointer pi - %v, *pi - %v\n", pi, &pi)
}

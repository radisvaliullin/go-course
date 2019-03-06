package main

import "fmt"

func main() {

	PrintInt(88)
}

func PrintInt(i int) {
	defer fmt.Println("defer i - ", i)

	i += 1

	fmt.Println("i - ", i)
}

package main

import (
	"fmt"
)

type MyError string

func (e MyError) Error() string {
	return "MyError"
}

func NilMyError() *MyError {
	return nil
}

func DoSome() (err error) {
	return NilMyError()
}

func main() {

	var err error
	fmt.Printf("err - %v\n", err)
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is Not nil")
	}

	err = DoSome()
	fmt.Printf("err - %v\n", err)
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is Not nil")
	}
}

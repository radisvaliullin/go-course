package main

import (
	"fmt"
	"io"
	"log"
)

func main() {

	bs := make([]*resp, 10)

	func() {
		for i := 0; i < 10; i++ {
			b := get()
			defer b.body.Close()
			bs[i] = b
		}
	}()

	for i := 0; i < 10; i++ {
		out := make([]byte, 10)
		n, err := bs[i].body.Read(out)
		if err != nil {
			log.Fatal("body read err - ", err)
		}
		fmt.Printf("i - %v; out - %v\n", i, string(out[:n]))
	}
	fmt.Printf("%v\n", bs)
}

type myReadCloser struct {
	msg string
}

func (rc *myReadCloser) Read(p []byte) (n int, err error) {
	n = copy(p, rc.msg)
	return n, nil
}

func (rc *myReadCloser) Close() error {
	rc.msg = "closed"
	return nil
}

type resp struct {
	body io.ReadCloser
}

func get() *resp {
	return &resp{body: &myReadCloser{msg: "msg"}}
}

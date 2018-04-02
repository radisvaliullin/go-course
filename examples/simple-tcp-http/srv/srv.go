package main

import (
	"errors"
	"log"
	"net"
	"time"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:7373")
	if err != nil {
		log.Fatalf("srv: listen: err - %v", err)
	}
	log.Println("srv: listen, begin.")

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Println("srv: accept: err - ", err)
			continue
		}

		log.Print("srv: new conn: addr - ", conn.RemoteAddr())
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	// read request
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		log.Printf("srv: conn - %v; err - %v; buff - %v",
			conn.RemoteAddr(), err, string(buff[:n]))
		return
	}
	log.Printf("srv: conn - %v; msg - \n%v\n", conn.RemoteAddr(), string(buff[:n]))

	// send response
	rawResp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: 6\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"Date: Wed, 19 Jul 1972 19:00:00 GMT\r\n\r\n" +
		"Hello.\n"
	resp := []byte(rawResp)

	// write response
	n, err = conn.Write(resp)
	if err != nil {
		log.Printf("srv: conn - %v; write resp err - %v", conn.RemoteAddr(), err)
		return
	} else if n != len(resp) {
		log.Printf("srv: conn - %v; write resp err - %v",
			conn.RemoteAddr(), errors.New("write resp n != len(writeData)"))
	}
	log.Printf("srv: conn - %v; write ok", conn.RemoteAddr())

	time.Sleep(time.Second * 2)
}

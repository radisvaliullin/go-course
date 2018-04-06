package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/radisvaliullin/go-course/examples/simple-self-tpc-protocol/models"
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

	clnAddr := fmt.Sprint(conn.RemoteAddr())

	for {
		// read packet header
		h := models.Header{}
		err := h.Read(conn)
		if err != nil {
			log.Printf("srv: client - %v; header read err - %v", clnAddr, err)
			return
		}

		// read packet data bytes
		dataBytes := make([]byte, int(h.Len))
		// not reason to check n
		_, err = io.ReadFull(conn, dataBytes)
		if err != nil {
			log.Printf("srv: client - %v; read packet data bytes, err - %v", clnAddr, err)
			return
		}

		// parse read data
		err = parseData(h, dataBytes, clnAddr)
		if err != nil {
			log.Printf("srv: client - %v; packet data parse, err - %v", clnAddr, err)
			return
		}

		// resp packet
		hB, err := h.ToBytes()
		if err != nil {
			log.Printf("srv: client - %v; get response packet, err - %v", clnAddr, err)
			return
		}

		// send resp packet
		_, err = conn.Write(hB)
		if err != nil {
			log.Printf("srv: client - %v; write response packet, err - %v", clnAddr, err)
			return
		}

		log.Printf(
			"srv: client - %v: req/res: read pack: head - %v; req data - %v",
			clnAddr, hB, dataBytes)

	}
}

func parseData(h models.Header, data []byte, clnAddr string) error {

	if h.Type == 0 {
		om := models.OneModel{}
		err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &om)
		if err != nil {
			log.Printf("srv: client - %v; packet data parse, type - %v, err - %v", clnAddr, h.Type, err)
			return err
		}

		log.Printf("srv: client - %v; get packet, type - %v, body - %+v", clnAddr, h.Type, om)

	} else if h.Type == 1 {
		jm := models.JSONModel{}
		err := json.Unmarshal(data, &jm)
		if err != nil {
			log.Printf("srv: client - %v; packet data parse, type - %v, err - %v", clnAddr, h.Type, err)
			return err
		}

		log.Printf("srv: client - %v; get packet, type - %v, body - %+v", clnAddr, h.Type, jm)

	}
	return nil
}

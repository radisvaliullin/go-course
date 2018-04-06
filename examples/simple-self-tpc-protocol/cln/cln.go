package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/radisvaliullin/go-course/examples/simple-self-tpc-protocol/models"
)

var (
	// ErrClientWRLen -
	ErrClientWRLen = errors.New("client: wrong read/write len")
	// ErrClientWrongResp -
	ErrClientWrongResp = errors.New("client: wrong response")
)

func main() {

	for i := 0; i < 10; i++ {
		go client()
	}

	for {
		log.Printf("heartbit")
		time.Sleep(time.Second * 1)
	}
}

//
func client() {

	srvAddr := "localhost:7373"

	conn, err := net.Dial("tcp", srvAddr)
	if err != nil {
		log.Printf("client - %v, dial err - %v", srvAddr, err)
		return
	}
	defer conn.Close()

	clnAddr := fmt.Sprint(conn.LocalAddr())

	for {

		// send first packet
		// send one model
		// not need goroutine
		err := sendOneModelPacket(conn, clnAddr)
		if err != nil {
			log.Printf("client - %v: send one model packet err - %v", clnAddr, err)
			return
		}

		// do pause for test
		time.Sleep(time.Millisecond * 100)

		// send second type packet
		// send json model
		// not need goroutine
		err = sendJSONModelPacket(conn, clnAddr)
		if err != nil {
			log.Printf("client - %v: send json model packet err - %v", clnAddr, err)
			return
		}

		// do pause for test
		time.Sleep(time.Millisecond * 100)
	}
}

//
func sendOneModelPacket(conn net.Conn, clnAddr string) error {

	// packet data
	om := models.OneModel{
		B:     255,
		Int64: 0x0A0B0C0D0E0F,
	}

	// get packet data bytes
	// 1 byte for B, 8 bytes for int64
	dataLen := 1 + 8
	dataBytes := make([]byte, dataLen)
	dataBytes[0] = om.B
	binary.LittleEndian.PutUint64(dataBytes[1:], uint64(om.Int64))

	// packet header
	h := models.Header{
		Len:  uint16(len(dataBytes)),
		Type: 0,
	}
	headBytes, err := h.ToBytes()
	if err != nil {
		log.Printf("client - %v: get head bytes err %v", clnAddr, err)
		return err
	}

	// full packet bytes
	packBytes := append(headBytes, dataBytes...)

	// send one model
	n, err := conn.Write(packBytes)
	if err != nil {
		log.Printf("client - %v: send packet bytes err - %v", clnAddr, err)
		return err
	} else if n != len(packBytes) {
		log.Printf("client - %v: send packet bytes err - %v", clnAddr, "send bytes len not equal return n")
		return ErrClientWRLen
	}

	// read response
	// resp bytes is 3 bytes
	resBytes := make([]byte, 3)

	// wrong way use simple read method
	// n, err := conn.Read(resBytes)
	_, err = io.ReadFull(conn, resBytes)
	if err != nil {
		log.Printf("client - %v: read response err - %v", clnAddr, err)
		return err
	}

	// respBytes must be equal request header
	// Compare return int val
	if bytes.Compare(headBytes, resBytes) != 0 {
		return ErrClientWrongResp
	}

	//log.Printf("client - %v: send one model: send pack - %v; resp - %v", clnAddr, packBytes, resBytes)

	return nil
}

//
func sendJSONModelPacket(conn net.Conn, clnAddr string) error {

	// packet data
	om := models.JSONModel{
		Name: "Use the go Luke.",
	}

	// get packet data bytes
	// marshal to json
	dataBytes, err := json.Marshal(&om)
	if err != nil {
		log.Printf("client - %v: get packet data bytes, err - %v", clnAddr, err)
		return err
	}

	// packet header
	h := models.Header{
		Len: uint16(len(dataBytes)),
		// JSON type is 1
		Type: 1,
	}
	headBytes, err := h.ToBytes()
	if err != nil {
		log.Printf("client - %v: get head bytes err %v", clnAddr, err)
		return err
	}

	// full packet bytes
	packBytes := append(headBytes, dataBytes...)

	//log.Printf("client - %v: send json model: send pack - %v", clnAddr, packBytes)

	// send one model
	n, err := conn.Write(packBytes)
	if err != nil {
		log.Printf("client - %v: send packet bytes err - %v", clnAddr, err)
		return err
	} else if n != len(packBytes) {
		log.Printf("client - %v: send packet bytes err - %v", clnAddr, "send bytes len not equal return n")
		return ErrClientWRLen
	}

	// read response
	// resp bytes is 3 bytes
	resBytes := make([]byte, 3)

	// wrong way use simple read method
	// n, err := conn.Read(resBytes)
	_, err = io.ReadFull(conn, resBytes)
	if err != nil {
		log.Printf("client - %v: json model: read response err - %v", clnAddr, err)
		return err
	}

	// respBytes must be equal request header
	// Compare return int val
	if bytes.Compare(headBytes, resBytes) != 0 {
		return ErrClientWrongResp
	}

	//log.Printf("client - %v: send json model: send pack - %v; resp - %v", clnAddr, packBytes, resBytes)

	return nil
}

package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"

	"github.com/radisvaliullin/go-course/examples/simple-self-tpc-protocol/models"
)

func main() {

}

//
func client(port int) {

	addr := ":" + strconv.Itoa(port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("client - %v, dial err - %v", addr, err)
		return
	}
	defer conn.Close()

	for {

		// send first package
		// send one model

		om := models.OneModel{
			B:     255,
			Int64: 0x0A0B0C0D0E0F,
		}

		// 1 byte for B, 8 bytes for int64
		dataLen := 1 + 8
		buf := make([]byte, dataLen)
		buf[0] = om.B
		binary.LittleEndian.PutUint64(buf[1:], uint64(om.Int64))

		h := models.Header
	}
}

package models

import (
	"bytes"
	"encoding/binary"
	"io"
)

// OneModel - for translate to bytes
type OneModel struct {
	B     byte
	Int64 int64
}

// JSONModel - for serialize to JSON
type JSONModel struct {
	Name string
}

// protocol models

// Header - self protocol packet header
// Byte order for header and body is little endian
type Header struct {
	// after header send data length
	Len uint16
	// 0 - for onemodel, 1 - for jsonmodel
	Type byte
}

// Read - read header from reader
func (h *Header) Read(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, h)
}

// ToBytes - get header bytes representation
func (h *Header) ToBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, h)
	return buf.Bytes(), err
}

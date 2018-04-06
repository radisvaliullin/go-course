package models

import (
	"bytes"
	"testing"
)

func TestHeader_Read(t *testing.T) {

	testBuff := []byte{1, 0, 1}

	h := Header{}

	err := h.Read(bytes.NewReader(testBuff))
	if err != nil {
		t.Error("header read err ", err)
		return
	}
	t.Log("head read - ", h)
}

func TestHeader_ToBytes(t *testing.T) {

	h := Header{1, 1}

	buf, err := h.ToBytes()
	if err != nil {
		t.Error("header to bytes err ", err)
		return
	}
	t.Log("head bytes - ", buf)
}

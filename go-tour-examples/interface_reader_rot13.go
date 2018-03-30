package main

import (
	"io"
	"os"
	"strings"
)

var rotTable = []byte("NOPQRSTUVWXYZABCDEFGHIJKLM[\\]^_`nopqrstuvwxyzabcdefghijklm")

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	rot13(b[:n])
	return n, err
}

func rot13(b []byte) {
	for i, ord := range b {
		// ASCII: A = 65, Z = 90, a = 97, z = 122
		if (ord >= 65 && ord <= 90) || (ord >= 97 && ord <= 122) {
			b[i] = rotTable[ord-65]
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

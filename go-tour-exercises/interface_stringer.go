package main

import "fmt"

type IPAddr [4]byte

//
func (a IPAddr) String() string {
	s := ""
	sep := ""
	for _, v := range a {
		s += sep
		s += fmt.Sprintf("%v", v)
		sep = "."
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

package main

import "fmt"

// Base type can has methods, base is not pointer and interface
type Some struct{}

// Method is func (String) with reciver (s Some).
func (s Some) String() string {
	return "Some is base type"
}

// // Type Can't has two method with value anp pointer reciver simultaneously.
// func (s *Some) String() string {
// 	return "Some is base type"
// }

//
func (s *Some) Error() string {
	return "Some Error"
}

func main() {

	i := 0

	s := Some{}
	sp := &s
	s.Error()
	sp.Error()
	s.String()
	sp.String()

	fmt.Printf("%v %s", i, s)
}

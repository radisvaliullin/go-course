package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	delta := 0.000000000001

	z := x
	for ; !(z*z-x < delta && z*z-x > -1.0*delta); z -= (z*z - x) / (2 * z) {

	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

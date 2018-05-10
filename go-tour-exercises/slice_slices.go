package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	sly := make([][]uint8, dy)

	for y := range sly {

		slx := make([]uint8, dx)

		for x := range slx {
			slx[x] = uint8((x + y) / 2)
		}
		sly[y] = slx
	}

	return sly
}

func main() {
	pic.Show(Pic)
}

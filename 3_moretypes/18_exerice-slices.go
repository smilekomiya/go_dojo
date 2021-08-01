package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		yarray := make([]uint8, dy)
		for y := 0; y < dy; y++ {
			yarray[y] = uint8(x * y)
		}
		result[x] = yarray
	}
	return result
}

func main() {
	pic.Show(Pic)
}

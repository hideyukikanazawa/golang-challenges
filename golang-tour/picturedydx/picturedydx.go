package main

import (
	"golang.org/x/tour/pic"
)

// Pic Function to generate dy, dx pixels
func Pic(dx, dy int) [][]uint8 {
	pixels := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pixels[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pixels[i][j] = uint8(i*j+(i+j))
		}
	}
	return pixels
}

func main() {
	pic.Show(Pic)
}

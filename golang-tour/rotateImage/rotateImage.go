package main

import (
	"fmt"
)

func prepend(x []int, y int) []int {
	z := append([]int{y}, x[:len(x)-1]...)
	return z
}

func rotateImage(input [][]int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}
	rows := len(input)
	columns := len(input[0])
	output := make([][]int, columns)
	for i := 0; i < columns; i++ {
		output[i] = make([]int, rows)
	}
	for _, row := range input {
		for j, value := range row {
			output[j] = prepend(output[j], value)
		}
	}
	return output
}


func rotateImageBetter(input [][]int) [][]int {
	n := len(input)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			input[i][j], input[j][i] = input[j][i], input[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for x, y := 0, n - 1; x < y; x, y = x+1, y-1 {
			input[i][x], input[i][y] = input[i][y], input[i][x]
		}
	}
	return input
}

func main() {
	imageI := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println("INITIAL")
	for _, v := range imageI {
		fmt.Println(v)
	}
	imageF := rotateImageBetter(imageI)
	fmt.Println("FINAL")
	for _, v := range imageF {
		fmt.Println(v)
	}
}

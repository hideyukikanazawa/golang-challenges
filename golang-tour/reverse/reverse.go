package main

import (
	"fmt"
)

func reverse(input string) string {
	output := []rune(input)
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return string(output)
}

func main() {
	s := "Hello Paul"
	t := reverse(s)
	fmt.Printf("Initial: %v\n", s)
	fmt.Printf("Final: %v", t)
}

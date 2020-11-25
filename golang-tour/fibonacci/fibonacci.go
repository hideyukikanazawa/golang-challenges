package main

import (
	"fmt"
)

func fibonacci(x int) int {
	i, j := 0, 1
	if x == 1 {
		return 0
	}
	for n := 1; n < x-1; n++ {
		i, j = j, i+j
	}
	return j
}

func main() {
	x := 4
	fib := fibonacci(x)
	fmt.Printf("Fibonacci number %v is %v", x, fib)
}

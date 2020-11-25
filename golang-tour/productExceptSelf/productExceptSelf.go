package main

import (
	"fmt"
)

func productExceptSelf(arr []int) []int {
	N := len(arr)
	leftArray, rightArray, outputArray := make([]int, N), make([]int, N), make([]int, N)
	leftArray[0], rightArray[N-1] = 1, 1
	for i := 1; i < N; i++ {
		leftArray[i] = arr[i-1] * leftArray[i-1]
	}
	for i := N - 2; i >= 0; i-- {
		rightArray[i] = arr[i+1] * rightArray[i+1]
	}
	for i := 0; i < N; i++ {
		outputArray[i] = leftArray[i] * rightArray[i]
	}
	return outputArray
}

func productExceptSelfOptimized(arr []int) []int {
    N := len(arr)
    outputArray := make([]int, N)
    outputArray[0] = 1
    for i := 1;i < N; i++ {
        outputArray[i] = outputArray[i-1] * arr[i-1]
    }
    R := 1
    for i := N-1; i >= 0; i-- {
        outputArray[i] *= R
        R *= arr[i]
    }
    return outputArray
}

func main() {
	initialArray := []int{1, 2, 3, 4}
	finalArray := productExceptSelfOptimized(initialArray)
	fmt.Printf("Initial array:\t%v\n", initialArray)
	fmt.Printf("Final array:\t%v", finalArray)
}

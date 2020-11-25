package main

import (
	"fmt"
)

func findLongestSubarrayBySum(arr []int, sum int) []int {
	N := len(arr)
	result := []int{-1}

	curSum, left, right := 0, 0, 0
	for right < N {
		curSum += arr[right]
		for left < right && curSum > sum {
			curSum -= arr[left]
			left++
		}
		if curSum == sum && (len(result) == 1 || result[1] - result[0] < right - left) {
			result = []int{left + 1, right + 1}
		}
		right++
	}
	return result
}

func main() {
	array := []int{1, 2, 3, 4, 5, 0, 0, 0, 6, 7, 8, 9, 10}
	sum := 21
	bounds := findLongestSubarrayBySum(array, sum)
	fmt.Printf("Boundaries are: %v", bounds)
}

package main

import (
	"fmt"
)

// Unit tests
func main() {
	// Test with an array containing both positive and negative numbers
	fmt.Println(MaxSubarraySum([]int{-5, 1, -3, 4, -9, 2, 1, -5, 4}))

	// Test with an array containing only negative numbers
	fmt.Println(MaxSubarraySum([]int{-2, -3, -1, -5}))

	// Test with an array containing only positive numbers
	fmt.Println(MaxSubarraySum([]int{1, 2, 3, 4}))

	// Test with an array containing a single element
	fmt.Println(MaxSubarraySum([]int{5}))
}

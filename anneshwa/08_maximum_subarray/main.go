package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize maxSum to the first element of the slice.
	maxSum := nums[0]
	// Initialize currentSum to the first element of the slice.
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// If the input slice is empty, return 0.
		if currentSum < 0 {
			// Start a new subarray
			currentSum = nums[i]
		} else {
			// Extend the current subarray
			currentSum += nums[i]
		}

		// If the currentSum is greater than maxSum, update maxSum.
		if currentSum > maxSum {
			// Update maxSum if we found a larger sum
			maxSum = currentSum
		}
	}
	// Return the maximum sum found.
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum sum of contiguous subarray:", maxSubArray(nums)) // Output: 6
}

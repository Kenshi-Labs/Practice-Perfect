package main

// MaxSubarraySum finds the contiguous subarray with the largest sum
func MaxSubarraySum(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for _, num := range nums[1:] {
		currentSum = max(num, currentSum+num)
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

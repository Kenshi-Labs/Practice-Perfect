package main

import (
	"math"
)

// FindMaxMin takes an array of integers and returns the maximum and minimum values.
func FindMaxMin(arr []int) (max int, min int) {
	// Initialize max to the smallest possible integer value and min to the largest
	max = math.MinInt64
	min = math.MaxInt64

	for _, v := range arr {
		if v > max {
			max = v // Update max if the current value is greater
		}
		if v < min {
			min = v // Update min if the current value is smaller
		}
	}
	return
}

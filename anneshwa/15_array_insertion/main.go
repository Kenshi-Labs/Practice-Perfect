package main

import (
	"fmt"
)

// Intersection returns an array of unique common elements between two arrays.
func Intersection(nums1, nums2 []int) []int {
	var result []int

	// Check each element in nums1
	for _, num1 := range nums1 {
		// Check if it is in nums2
		for _, num2 := range nums2 {
			if num1 == num2 {
				// Check if it is not already added to result
				if !contains(result, num1) {
					result = append(result, num1)
				}
			}
		}
	}

	return result
}

// contains checks if a number is in the result slice.
func contains(slice []int, num int) bool {
	for _, n := range slice {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{0, 8, 9}
	result := Intersection(nums1, nums2)
	fmt.Println(result) // Output: [2]
}

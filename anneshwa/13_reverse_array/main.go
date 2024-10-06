package main

import (
	"errors"
	"fmt"
)

// Reverse reverses the elements of the array from start to end indices.
func Reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Rotate rotates the array to the right by k steps.
func Rotate(nums []int, k int) ([]int, error) {
	n := len(nums)
	if n == 0 {
		return nil, errors.New("array is empty")
	}
	if k <= 0 {
		return nil, errors.New("rotation steps must be greater than 0")
	} else if k > n {
		return nums, nil
	}

	k = k % n // Normalize k to be within the bounds of the array length

	if k == 0 {
		return nums, nil // No need to rotate if k is 0 after normalization
	}

	// Reverse the entire array
	Reverse(nums, 0, n-1)

	// Reverse the first k elements
	Reverse(nums, 0, k-1)

	// Reverse the remaining n-k elements
	Reverse(nums, k, n-1)

	return nums, nil
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 70

	// Capture the returned values
	rotated, err := Rotate(nums, k)
	if err != nil {
		// Print the error if one occurred
		fmt.Println("Error:", err)
	} else {
		// Print the rotated array
		fmt.Println(rotated) // Output: [4, 5, 1, 2, 3]
	}
}

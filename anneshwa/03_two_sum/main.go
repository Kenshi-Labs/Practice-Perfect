package main

import (
	"fmt"

)

func main() {
	
	// Testing TwoSum
	fmt.Println("Testing TwoSum:")
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9)) // Output: [0 1]
	fmt.Println(TwoSum([]int{3, 2, 4}, 6))      // Output: [1 2]
	fmt.Println()
}


// PROBLEM 3: TwoSum finds indices of two numbers that add up to a target
func TwoSum(nums []int, target int) []int {
	// Iterate over each number in the array
	for i := 0; i < len(nums); i++ {
		// Check for each number, if there is another number that sums up to the target
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				// Return the indices of the two numbers
				return []int{i, j}
			}
		}
	}

	// If no such pair is found, return an empty slice
	return nil
}

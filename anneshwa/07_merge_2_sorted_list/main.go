package main

import (
	"fmt"
)

func mergeTwoLists(l1 []int, l2 []int) []int {
	// Initialize an empty slice to store the merged list
	result := []int{}
	i, j := 0, 0

	// Iterate through both lists until either list is fully traversed
	for i < len(l1) && j < len(l2) {
		if l1[i] <= l2[j] {
			// Append current element of l1 to result & move pointer in l1 to the next element
			result = append(result, l1[i])
			i++
		} else {
			result = append(result, l2[j])
			j++
		}
	}

	// Append any remaining elements in l1 to the result slice
	for i < len(l1) {
		result = append(result, l1[i])
		i++
	}

	// Append any remaining elements in l2 to the result slice
	for j < len(l2) {
		result = append(result, l2[j])
		j++
	}
	// Return the merged and sorted slice
	return result
}

func main() {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}

	merged := mergeTwoLists(l1, l2)
	fmt.Println(merged) // Output: [1 1 2 3 4 4]
}

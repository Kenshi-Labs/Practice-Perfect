package main

// removeDuplicates removes duplicates from a sorted array nums in-place.
// It returns the new length of the array after removal.
func removeDuplicates(nums []int) int {
	// If the array is empty, there are no elements to remove duplicates from.
	if len(nums) == 0 {
		return 0
	}

	// slow-runner pointer
	slowPtr := 0

	for genPtr := 1; genPtr < len(nums); genPtr++ {
		// Compare current element (nums[genPtr]) with previous distinct element (nums[slowPtr])
		if nums[genPtr] != nums[slowPtr] {
			// Move slowPtr to the next position
			slowPtr++
			nums[slowPtr] = nums[genPtr] // Copy the unique element to nums[slowPtr]
		}
	}

	// Length of the array with duplicates removed
	return slowPtr + 1
}

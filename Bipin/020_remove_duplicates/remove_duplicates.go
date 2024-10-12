package main

// RemoveDuplicates removes duplicates in-place from a sorted array and returns the new length.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize two pointers: `i` for the unique elements
	i := 0

	// Iterate through the array starting from the second element
	for j := 1; j < len(nums); j++ {
		// If the current element is different from the previous unique element
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	// Return the length of the array with unique elements
	return i + 1
}

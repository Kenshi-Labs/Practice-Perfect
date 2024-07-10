package main

// twoSum finds two numbers in nums that add up to target and returns their indices.
func twoSum(nums []int, target int) []int {
	// Create a map to store the complement of each number and its index.
	numMap := make(map[int]int)

	// Iterate through the array.
	for i, num := range nums {
		// Calculate the complement needed to reach the target.
		complement := target - num

		// Check if the complement exists in the map.
		if j, found := numMap[complement]; found {
			// Return the indices of the two numbers.
			return []int{j, i}
		}

		// Store the current number and its index in the map.
		numMap[num] = i
	}

	// If no solution is found, return an empty slice (though by problem statement there will always be a solution).
	return []int{}
}

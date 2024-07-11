package main

import (
	"testing"
)

// Define a struct for test cases
type testCase struct {
	nums        []int
	expectedLen int
	expectedArr []int
	description string
}

func TestRemoveDuplicates(t *testing.T) {
	// Define test cases using the struct
	testCases := []testCase{
		{
			nums:        []int{1, 1, 2},
			expectedLen: 2,
			expectedArr: []int{1, 2},
			description: "Array with duplicates",
		},
		{
			nums:        []int{1, 2, 3},
			expectedLen: 3,
			expectedArr: []int{1, 2, 3},
			description: "Array with no duplicates",
		},
		{
			nums:        []int{},
			expectedLen: 0,
			expectedArr: []int{},
			description: "Empty array",
		},
		{
			nums:        []int{1, 1, 1, 1, 1},
			expectedLen: 1,
			expectedArr: []int{1},
			description: "Array with all identical elements",
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Make a copy of nums for testing
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums)

			// Call removeDuplicates
			length := removeDuplicates(numsCopy)

			// Check length
			if length != tc.expectedLen {
				t.Errorf("Expected length %d, but got %d", tc.expectedLen, length)
			}

			// Check resulting array
			for i := 0; i < length; i++ {
				if numsCopy[i] != tc.expectedArr[i] {
					t.Errorf("Expected array %v, but got %v", tc.expectedArr, numsCopy[:length])
					break
				}
			}
		})
	}
}

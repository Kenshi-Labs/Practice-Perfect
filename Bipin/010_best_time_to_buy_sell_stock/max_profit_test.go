package main

import (
	"testing"
)

// TestMaxProfit contains all the unit tests for the MaxProfit function
func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5}, // Max profit is 6 - 1 = 5
		{[]int{7, 6, 4, 3, 1}, 0},    // No profit is possible in a descending array
		{[]int{2, 4, 1, 4, 6, 2}, 5}, // Profit from 1 to 6, max profit is 6 - 1 = 5
		{[]int{1, 1}, 0},             // Only two elements, no profit can be made
	}

	for _, test := range tests {
		result := MaxProfit(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}

package main

import (
	"testing"
)

// Unit test for FindMaxMin
func TestFindMaxMin(t *testing.T) {
	// 1. Test with an array of positive numbers
	t.Run("Positive numbers", func(t *testing.T) {
		input := []int{10, 5, 8, 2, 15}
		expectedMax := 15
		expectedMin := 2
		max, min := FindMaxMin(input)
		if max != expectedMax || min != expectedMin {
			t.Errorf("For input %v, expected Max = %d, Min = %d; got Max = %d, Min = %d", input, expectedMax, expectedMin, max, min)
		}
	})

	// 2. Test with an array containing negative numbers
	t.Run("Negative numbers", func(t *testing.T) {
		input := []int{-10, -5, -8, -2, -15}
		expectedMax := -2
		expectedMin := -15
		max, min := FindMaxMin(input)
		if max != expectedMax || min != expectedMin {
			t.Errorf("For input %v, expected Max = %d, Min = %d; got Max = %d, Min = %d", input, expectedMax, expectedMin, max, min)
		}
	})

	// 3. Test with an array containing duplicate values
	t.Run("Duplicate values", func(t *testing.T) {
		input := []int{5, 3, 5, 3, 5}
		expectedMax := 5
		expectedMin := 3
		max, min := FindMaxMin(input)
		if max != expectedMax || min != expectedMin {
			t.Errorf("For input %v, expected Max = %d, Min = %d; got Max = %d, Min = %d", input, expectedMax, expectedMin, max, min)
		}
	})

	// 4. Test with an array of length 1
	t.Run("Single element", func(t *testing.T) {
		input := []int{42}
		expectedMax := 42
		expectedMin := 42
		max, min := FindMaxMin(input)
		if max != expectedMax || min != expectedMin {
			t.Errorf("For input %v, expected Max = %d, Min = %d; got Max = %d, Min = %d", input, expectedMax, expectedMin, max, min)
		}
	})
}

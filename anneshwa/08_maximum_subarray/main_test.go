package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	// Test case 1: Test with an array containing both positive and negative numbers
	t.Run("Mixed numbers", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		expected := 6
		assert.Equal(t, expected, maxSubArray(nums), "Expected and actual results do not match")
	})

	// Test case 2: Test with an array containing only negative numbers
	t.Run("All negative numbers", func(t *testing.T) {
		nums := []int{-2, -3, -1, -5, -4}
		expected := -1
		assert.Equal(t, expected, maxSubArray(nums), "Expected and actual results do not match")
	})

	// Test case 3: Test with an array containing only positive numbers
	t.Run("All positive numbers", func(t *testing.T) {
		nums := []int{2, 3, 1, 5, 4}
		expected := 15
		assert.Equal(t, expected, maxSubArray(nums), "Expected and actual results do not match")
	})

	// Test case 4: Test with an array containing a single element
	t.Run("Single element", func(t *testing.T) {
		nums := []int{7}
		expected := 7
		assert.Equal(t, expected, maxSubArray(nums), "Expected and actual results do not match")
	})
}

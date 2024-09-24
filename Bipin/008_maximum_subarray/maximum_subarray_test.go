package main

import (
	"testing"
)

func TestMaxSubarraySum(t *testing.T) {
	// Test with an array containing both positive and negative numbers
	result := MaxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 9, -5, 4})
	expected := 14
	if result != expected {
		t.Errorf("Test failed, expected %d, got %d", expected, result)
	}

	// Test with an array containing only negative numbers
	result = MaxSubarraySum([]int{-2, -3, -1, -5})
	expected = -1
	if result != expected {
		t.Errorf("Test failed, expected %d, got %d", expected, result)
	}

	// Test with an array containing only positive numbers
	result = MaxSubarraySum([]int{1, 2, 3, 4, 0})
	expected = 10
	if result != expected {
		t.Errorf("Test failed, expected %d, got %d", expected, result)
	}

	// Test with an array containing a single element
	result = MaxSubarraySum([]int{7})
	expected = 7
	if result != expected {
		t.Errorf("Test failed, expected %d, got %d", expected, result)
	}

	// Test with an array containing alternating negative and positive numbers
	result = MaxSubarraySum([]int{-1, 2, -1, 2, -1, 2})
	expected = 4
	if result != expected {
		t.Errorf("Test failed, expected %d, got %d", expected, result)
	}
}

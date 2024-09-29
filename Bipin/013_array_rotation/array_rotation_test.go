package main

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	// Test case 1: Non-empty array with positive rotation
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	RotateArray(nums, k)
	expected := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("Test Case 1 Failed: expected %v, got %v", expected, nums)
	}

	// Test case 2: Rotation equal to array length
	nums = []int{1, 2, 3, 4, 5}
	k = 5
	RotateArray(nums, k)
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("Test Case 2 Failed: expected %v, got %v", expected, nums)
	}

	// Test case 3: Rotation greater than array length
	nums = []int{1, 2, 3, 4, 5}
	k = 7
	RotateArray(nums, k)
	expected = []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("Test Case 3 Failed: expected %v, got %v", expected, nums)
	}

	// Test case 4: Empty array
	nums = []int{}
	k = 3
	RotateArray(nums, k)
	expected = []int{}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("Test Case 4 Failed: expected %v, got %v", expected, nums)
	}
}

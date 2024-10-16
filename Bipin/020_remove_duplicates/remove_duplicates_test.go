package main

import (
	"reflect"
	"testing"
)

// TestRemoveDuplicates contains unit tests for the removeDuplicates function.
func TestRemoveDuplicates(t *testing.T) {
	// 1. Test with an array containing duplicates
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedNums1 := []int{0, 1, 2, 3, 4}
	expectedLength1 := 5
	length1 := RemoveDuplicates(nums1)
	if length1 != expectedLength1 || !reflect.DeepEqual(nums1[:length1], expectedNums1) {
		t.Errorf("Expected length %d, nums %v, but got length %d, nums %v", expectedLength1, expectedNums1, length1, nums1[:length1])
	}

	// 2. Test with an array with no duplicates
	nums2 := []int{1, 2, 3, 4, 5}
	expectedNums2 := []int{1, 2, 3, 4, 5}
	expectedLength2 := 5
	length2 := RemoveDuplicates(nums2)
	if length2 != expectedLength2 || !reflect.DeepEqual(nums2[:length2], expectedNums2) {
		t.Errorf("Expected length %d, nums %v, but got length %d, nums %v", expectedLength2, expectedNums2, length2, nums2[:length2])
	}

	// 3. Test with an empty array
	nums3 := []int{}
	expectedNums3 := []int{}
	expectedLength3 := 0
	length3 := RemoveDuplicates(nums3)
	if length3 != expectedLength3 || !reflect.DeepEqual(nums3[:length3], expectedNums3) {
		t.Errorf("Expected length %d, nums %v, but got length %d, nums %v", expectedLength3, expectedNums3, length3, nums3[:length3])
	}

	// 4. Test with an array containing all identical elements
	nums4 := []int{1, 1, 1, 1, 1}
	expectedNums4 := []int{1}
	expectedLength4 := 1
	length4 := RemoveDuplicates(nums4)
	if length4 != expectedLength4 || !reflect.DeepEqual(nums4[:length4], expectedNums4) {
		t.Errorf("Expected length %d, nums %v, but got length %d, nums %v", expectedLength4, expectedNums4, length4, nums4[:length4])
	}
}

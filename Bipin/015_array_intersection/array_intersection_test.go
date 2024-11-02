package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	// Test 1: Arrays with common elements
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2, 3}
	expected := []int{2}
	result := Intersection(nums1, nums2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test 1 failed. Expected %v, got %v", expected, result)
	}

	// Test 2: Arrays with no common elements
	nums1 = []int{4, 5, 6}
	nums2 = []int{7, 8, 9}
	expected = []int{}
	result = Intersection(nums1, nums2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test 2 failed. Expected %v, got %v", expected, result)
	}

	// Test 3: Both arrays are empty
	nums1 = []int{}
	nums2 = []int{}
	expected = []int{}
	result = Intersection(nums1, nums2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test 3 failed. Expected %v, got %v", expected, result)
	}

	// Test 4: Arrays with duplicates
	nums1 = []int{1, 1, 2, 2}
	nums2 = []int{2, 2, 1, 1}
	expected = []int{2, 1}
	result = Intersection(nums1, nums2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test 4 failed. Expected %v, got %v", expected, result)
	}
}

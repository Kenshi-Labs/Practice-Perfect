package main

import (
	"fmt"
	"testing"
)

func TestFind_a_sum(t *testing.T) {
	testcases := []struct {
		array    []int
		target   int
		expected []int
	}{
		{array: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{array: []int{2, 7, 11, 15}, target: 120, expected: []int{}},
		{array: []int{20, 7, 2, 15}, target: 22, expected: []int{0, 2}},
		{array: []int{1, 8}, target: 9, expected: []int{0, 1}},
	}

	for i, val := range testcases {
		output := Find_a_sum(val.array, val.target)
		if !CompareTwoArray(output, val.expected) {
			t.Errorf("Error in output for the testcase no %d for input of %d", i+1, val.array)
		} else {
			fmt.Printf("Test case %d passed \n", i+1)
		}
	}
}

func CompareTwoArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

package main

import (
	"fmt"
	"testing"
)

func TestMaximum_subarray(t *testing.T) {
	testcases := []struct{
		input []int
		expected int
	}{
		{input : []int{-2,1,-3,4,-1,2,1,-5,4},expected: 6}, // containing both positive and negative numbers
		{input : []int{-2,-1,-3,-5},expected: -1}, // containing only negative numbers
		{input : []int{5,6,7,8},expected: 26}, // containing only positive numbers
		{input : []int{-1},expected: -1}, // containing a single element
	}

	for i, val := range testcases {
		if maxSubArray(val.input) != val.expected {
			t.Errorf("Error in the testcase no. %d for the input of %v", i+1, val.input)
		}else{
			fmt.Printf("Test case %d passed \n", i+1)
		}
	}
}
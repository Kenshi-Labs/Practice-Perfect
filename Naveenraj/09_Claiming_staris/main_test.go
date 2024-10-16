package main

import (
	"fmt"
	"testing"
)

func TestClaimingStaris(t *testing.T) {
	testcases := []struct {
		arr      int
		expected int
	}{
		{arr: 1, expected: 1},
		{arr: 2, expected: 2},
		{arr: 5, expected: 8},
		{arr: 0, expected: 0},
	}
	for i, val := range testcases {
		output := climbStairs(val.arr)

		if output != val.expected {
			t.Errorf("Error in the test case no %d for the input of %d", i+1, val.arr)
		} else {
			fmt.Printf("Test case %d passes \n", i+1)
		}
	}
}

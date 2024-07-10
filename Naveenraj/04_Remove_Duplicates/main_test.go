package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplic(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{arr: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: 5},
		{arr: []int{0, 1, 2, 3, 4}, expected: 5},
		{arr: []int{}, expected: 0},
		{arr: []int{0, 0, 0, 0, 0}, expected: 1},
	}

	for i, val := range testcases {
		output := Remove_Duplicates(val.arr)

		if output != val.expected {
			t.Errorf("Error in the test case no %d for the input of %d", i+1, val.arr)
		} else {
			fmt.Printf("Test case %d passes \n", i+1)
		}
	}
}

package main

import (
	"fmt"
	"testing"
)

func TestBestTimeToBuySellStock(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{arr: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{arr: []int{7, 6, 4, 3, 1}, expected: 0},
	}
	for i, val := range testcases {
		output := maxProfit(val.arr)

		if output != val.expected {
			t.Errorf("Error in the test case no %d for the input of %d", i+1, val.arr)
		} else {
			fmt.Printf("Test case %d passes \n", i+1)
		}
	}
}

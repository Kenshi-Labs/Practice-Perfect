package main

import (
	"fmt"
	"testing"
)

func Test_Longest_common_prefix(t *testing.T) {
	testcases := []struct {
		arr      []string
		expected string
	}{
		{arr: []string{"flower", "flow", "flight"}, expected: "fl"},
		{arr: []string{"no", "common", "prefix"}, expected: ""},
		{arr: []string{"flower", "flow", ""}, expected: ""},
		{arr: []string{"flower"}, expected: "flower"},
	}

	for i, val := range testcases {
		output := Longest_common_prefix(val.arr)
		if output != val.expected {
			t.Errorf("Error in the testcase no. %d for the input of %v", i+1, val.arr)
		} else {
			fmt.Printf("Test case %d passed \n", i+1)
		}
	}
}

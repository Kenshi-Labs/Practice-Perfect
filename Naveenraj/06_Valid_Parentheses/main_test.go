package main

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	testcases := []struct {
		input          string
		expectedOutput bool
	}{
		{input: "()", expectedOutput: true},    // valid parentheses
		{input: "([})", expectedOutput: false}, // invalid parentheses
		{input: "", expectedOutput: false},     // empty string
		{input: "({[", expectedOutput: false},  // only opening brackets
	}

	for i, val := range testcases {
		output := Valid_Parentheses(val.input)
		if output != val.expectedOutput {
			t.Errorf("Error in the testcase no %d of input %s", i+1, val.input)
		} else {
			fmt.Printf("Test case %d passed \n", i+1)
		}
	}
}

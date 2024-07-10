package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"G o", "o G"},
		{"madam", "madam"},
	}

	for i, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.input, result, test.expected)
		} else {
			fmt.Printf("Test Case %d Pass \n", i+1)
		}
	}
}

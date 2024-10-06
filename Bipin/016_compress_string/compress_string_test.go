package main

import (
	"testing"
)

// compressString function (insert your method here)

// TestCompressString runs a variety of test cases using the testing package.
func TestCompressString(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"}, // Basic compression
		{"abcdef", "abcdef"},        // No compression needed
		{"", ""},                    // Empty string
		{"a", "a"},                  // Single character
		{"abcd", "abcd"},            // No repeated characters
	}

	// Iterate through test cases
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := CompressString(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, result)
			}
		})
	}
}

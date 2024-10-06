package main

import "testing"

// TestFirstNonRepeatingCharacter tests the FirstNonRepeatingCharacter function.
func TestFirstNonRepeatingCharacter(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    string
		expected rune
	}{
		{"aabbcdeeff", 'c'}, // Test with a string containing a non-repeating character
		{"aabbcc", 0},       // Test with a string where all characters repeat
		{"", 0},             // Test with an empty string
		{"x", 'x'},          // Test with a string containing only one character
	}

	// Iterate through each test case
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := FirstNonRepeatingCharacter(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected '%c' but got '%c'", test.input, test.expected, result)
			}
		})
	}
}

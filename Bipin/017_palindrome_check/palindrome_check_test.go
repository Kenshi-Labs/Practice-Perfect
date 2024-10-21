package main

import "testing"

// TestIsPalindrome tests the IsPalindrome function.
func TestIsPalindrome(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true}, // Test with a valid palindrome containing spaces and punctuation
		{"race a car", false},                    // Test with a non-palindrome
		{"", true},                               // Test with an empty string
		{"!!@@##$$", true},                       // Test with a string containing only non-alphanumeric characters
	}

	// Iterate through each test case
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsPalindrome(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected %t but got %t", test.input, test.expected, result)
			}
		})
	}
}

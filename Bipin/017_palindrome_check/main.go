package main

import "fmt"

// Unit Test Function to test the palindrome function
func testPalindrome() {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true}, // Test with a valid palindrome containing spaces and punctuation
		{"race a car", false},                    // Test with a non-palindrome
		{"", true},                               // Test with an empty string
		{"!!@@##$$", true},                       // Test with a string containing only non-alphanumeric characters
	}

	// Iterate through the test cases and compare results
	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			fmt.Printf("Test Failed: input='%s', expected=%t, got=%t\n", test.input, test.expected, result)
		} else {
			fmt.Printf("Test Passed: input='%s', expected=%t, got=%t\n", test.input, test.expected, result)
		}
	}
}

func main() {
	// Run the unit tests
	testPalindrome()
}

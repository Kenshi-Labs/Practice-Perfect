package main

import (
	"fmt"
)

// Function to find the first non-repeating character
func FirstNonRepeatingChar(s string) byte {
	// Create a map to store character frequencies
	charFrequency := make(map[byte]int)

	// Iterate over the string and populate the frequency map
	for i := 0; i < len(s); i++ {
		charFrequency[s[i]]++
	}

	// Iterate over the string again to find the first non-repeating character
	for i := 0; i < len(s); i++ {
		if charFrequency[s[i]] == 1 {
			return s[i]
		}
	}

	// If no non-repeating character is found, return a space (or any other default value)
	return ' '
}

func main() {
	// Test cases
	fmt.Println("Input: 'aabbcdeeff' | Output:", string(FirstNonRepeatingChar("aabbcdeeff"))) // Output: 'c'
	fmt.Println("Input: 'aabbcc' | Output:", string(FirstNonRepeatingChar("aabbcc")))         // Output: ' ' (no non-repeating character)
	fmt.Println("Input: '' | Output:", string(FirstNonRepeatingChar("")))                     // Output: ' ' (empty string)
	fmt.Println("Input: 'z' | Output:", string(FirstNonRepeatingChar("z")))                   // Output: 'z' (one character string)
}

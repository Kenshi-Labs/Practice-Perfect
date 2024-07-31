package main

import (
	"fmt"
	"strings"
)

// countVowels returns the number of vowels in the input string.
// It considers 'a', 'e', 'i', 'o', 'u' as vowels and is case-insensitive.
func countVowels(s string) int {
	// Define the set of vowels
	vowels := "aeiou"

	// Convert the input string to lowercase
	lowercaseStr := strings.ToLower(s)

	// Initialize a counter for vowels
	count := 0

	// Iterate through each character in the string
	for _, char := range lowercaseStr {
		// Check if the character is a vowel
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	// Return the total count of vowels
	return count
}

func main() {
	// Example usage
	input := "Hello, World!"
	vowelCount := countVowels(input)
	fmt.Printf("The number of vowels in \"%s\" is: %d\n", input, vowelCount)
}
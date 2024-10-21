package main

import "fmt"

// Test cases
func main() {
	// Test case 1: String with vowels
	fmt.Println("Test Case 1 - 'Hello, World!':", CountVowels("Hello, World!"))

	// Test case 2: String with no vowels
	fmt.Println("Test Case 2 - 'xyz123':", CountVowels("xyz123"))

	// Test case 3: Empty string
	fmt.Println("Test Case 3 - '':", CountVowels(""))

	// Test case 4: String with only vowels
	fmt.Println("Test Case 4 - 'aeiouAEIOU':", CountVowels("aeiouAEIOU"))
}

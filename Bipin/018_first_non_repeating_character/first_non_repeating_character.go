package main

// FirstNonRepeatingCharacter finds the first non-repeating character in a string
func FirstNonRepeatingCharacter(s string) rune {
	const alphabetSize = 256           // Assuming extended ASCII characters
	count := make([]int, alphabetSize) // To count occurrences of each character

	// First pass: Count occurrences of each character
	for _, char := range s {
		count[char]++
	}

	// Second pass: Find the first character that has a count of 1
	for _, char := range s {
		if count[char] == 1 {
			return char
		}
	}

	return 0 // Return 0 if there is no non-repeating character
}

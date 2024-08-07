package main

// reverseString reverses a given input string and returns the reversed string.
func reverseString(input string) string {
	// Convert the input string to a slice of runes.
	runes := []rune(input)

	// Use a single loop to swap characters from both ends towards the center.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap characters at position i and j.
	}

	// Convert the reversed slice of runes back to a string and return it.
	return string(runes)
}

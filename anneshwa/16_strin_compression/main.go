package main

import (
	"fmt"
)

// Compress compresses a string by counting consecutive characters.
func Compress(s string) string {
	if len(s) == 0 {
		return s // Return empty string if input is empty
	}

	var compressed string
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++ // Increase count for repeated characters
		} else {
			compressed += string(s[i-1]) + fmt.Sprint(count) // Append char and count
			count = 1 // Reset count for new character
		}
	}

	// Append the last character and its count
	compressed += string(s[len(s)-1]) + fmt.Sprint(count)

	// Return the original string if compressed string is not smaller
	if len(compressed) >= len(s) {
		return s
	}
	return compressed
}

func main() {
	input := "aabcccccaaa"
	output := Compress(input)
	fmt.Println(output) // Output: "a2b1c5a3"
}

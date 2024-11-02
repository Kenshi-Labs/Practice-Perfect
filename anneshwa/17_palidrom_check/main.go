package main

import (
	"fmt"
	"strings"
)

// IsPalindrome checks if a string is a palindrome, ignoring case, spaces, and punctuation.
func IsPalindrome(s string) bool {
	// Convert string to lowercase
	s = strings.ToLower(s)

	// Two-pointer approach to check for palindrome
	left, right := 0, len(s)-1
	for left < right {
		// Move left pointer to the next alphanumeric character
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Move right pointer to the previous alphanumeric character
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		// Check if characters are equal
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// isAlphanumeric checks if the given byte is a letter or a digit
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func main() {
	input := "A man, a plan, a canal: Panama"
	fmt.Println(IsPalindrome(input)) // Output: true
}

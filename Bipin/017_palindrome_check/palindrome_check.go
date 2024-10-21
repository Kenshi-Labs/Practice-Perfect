package main

import (
	"unicode"
)

// IsPalindrome checks if a given string is a palindrome, ignoring non-alphanumeric characters and case.
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Move the left pointer to the next alphanumeric character
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		// Move the right pointer to the previous alphanumeric character
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		// If characters don't match, it's not a palindrome
		if toLower(rune(s[left])) != toLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

// isAlphanumeric checks if the character is a letter or a number
func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// toLower converts a rune to lowercase (for case-insensitive comparison)
func toLower(r rune) rune {
	return unicode.ToLower(r)
}

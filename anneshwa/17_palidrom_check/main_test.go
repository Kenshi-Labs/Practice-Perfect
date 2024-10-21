package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	// Positive Test 1: Valid palindrome with spaces and punctuation
	t.Run("Valid Palindrome", func(t *testing.T) {
		input := "A man, a plan, a canal: Panama"
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for a valid palindrome")
	})

	// Negative Test 2: Non-palindrome
	t.Run("Non-palindrome", func(t *testing.T) {
		input := "This is not a palindrome"
		result := IsPalindrome(input)
		assert.False(t, result, "Expected false for a non-palindrome")
	})

	// Edge Case: Almost a palindrome, but one character is different
	t.Run("Almost Palindrome", func(t *testing.T) {
		input := "A man, a plan, a cannal: Panama" // 'cannal' has an extra 'n'
		result := IsPalindrome(input)
		assert.False(t, result, "Expected false for a string that's almost a palindrome but not quite")
	})

	// Edge Case: Palindrome with mixed case and non-alphanumeric characters
	t.Run("Mixed Case Palindrome", func(t *testing.T) {
		input := "No lemon, no melon!"
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for a mixed-case palindrome with non-alphanumeric characters")
	})

	// Edge Case: Empty string
	t.Run("Empty String", func(t *testing.T) {
		input := ""
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for an empty string")
	})

	// Edge Case: String with only non-alphanumeric characters
	t.Run("Non-alphanumeric characters only", func(t *testing.T) {
		input := "!@#$%^&*()"
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for a string with only non-alphanumeric characters")
	})

	// Negative Case: One character string that is not alphanumeric
	t.Run("Single Non-alphanumeric Character", func(t *testing.T) {
		input := "!"
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for a single non-alphanumeric character")
	})

	// Negative Case: String with non-palindromic alphanumeric characters
	t.Run("Non-palindromic Alphanumeric", func(t *testing.T) {
		input := "abcde"
		result := IsPalindrome(input)
		assert.False(t, result, "Expected false for a non-palindromic alphanumeric string")
	})

	// Negative Case: Special characters only but different ones
	t.Run("Different Special Characters", func(t *testing.T) {
		input := "!@#$%^"
		result := IsPalindrome(input)
		assert.True(t, result, "Expected true for different non-alphanumeric characters (ignoring them)")
	})
}

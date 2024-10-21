package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the firstNonRepeatingChar function
func TestFirstNonRepeatingChar(t *testing.T) {
	// Positive Test Case 1: A string with a non-repeating character
	t.Run("Non-Repeating Character", func(t *testing.T) {
		input := "aabbcdeeff"
		expected := byte('c') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected 'c' for input 'aabbcdeeff'")
	})

	// Positive Test Case 2: A string where all characters repeat
	t.Run("All Repeating Characters", func(t *testing.T) {
		input := "aabbcc"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input 'aabbcc'")
	})

	// Positive Test Case 3: An empty string
	t.Run("Empty String", func(t *testing.T) {
		input := ""
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for empty input")
	})

	// Positive Test Case 4: A string with only one character
	t.Run("Single Character String", func(t *testing.T) {
		input := "z"
		expected := byte('z') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected 'z' for input 'z'")
	})

	// Positive Test Case 5: A string with non-repeating character at the end
	t.Run("Non-Repeating at End", func(t *testing.T) {
		input := "aabbccd"
		expected := byte('d') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected 'd' for input 'aabbccd'")
	})

	// Negative Test Case 1: All characters are uppercase but repeat
	t.Run("All Uppercase Repeating", func(t *testing.T) {
		input := "AABBCC"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input 'AABBCC'")
	})

	// Negative Test Case 2: String with numbers where all characters repeat
	t.Run("Numbers All Repeating", func(t *testing.T) {
		input := "112233"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input '112233'")
	})

	// Negative Test Case 3: Single repeating character
	t.Run("Single Repeating Character", func(t *testing.T) {
		input := "aa"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input 'aa'")
	})

	// Negative Test Case 4: String with special characters where all repeat
	t.Run("Special Characters Repeating", func(t *testing.T) {
		input := "!!@@##"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input '!!@@##'")
	})

	// Negative Test Case 5: String with a mix of letters and numbers where all characters repeat
	t.Run("Mixed Letters and Numbers Repeating", func(t *testing.T) {
		input := "a1a1b2b2"
		expected := byte(' ') // Use byte instead of rune
		result := FirstNonRepeatingChar(input)
		assert.Equal(t, expected, result, "Expected ' ' for input 'a1a1b2b2'")
	})
}

package main

import (
	"testing"
)

// Unit test for isValid
func TestIsValid(t *testing.T) {
	// 1. Test with a valid string of parentheses
	t.Run("Valid parentheses", func(t *testing.T) {
		input := "()[]{}"
		expected := true
		result := IsValid(input)
		if result != expected {
			t.Errorf("For input '%s', expected %v; got %v", input, expected, result)
		}
	})

	// 2. Test with an invalid string of parentheses
	t.Run("Invalid parentheses", func(t *testing.T) {
		input := "(]"
		expected := false
		result := IsValid(input)
		if result != expected {
			t.Errorf("For input '%s', expected %v; got %v", input, expected, result)
		}
	})

	// 3. Test with an empty string
	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := true // An empty string is considered valid
		result := IsValid(input)
		if result != expected {
			t.Errorf("For input '%s', expected %v; got %v", input, expected, result)
		}
	})

	// 4. Test with a string containing only opening brackets
	t.Run("Only opening brackets", func(t *testing.T) {
		input := "((("
		expected := false
		result := IsValid(input)
		if result != expected {
			t.Errorf("For input '%s', expected %v; got %v", input, expected, result)
		}
	})
}

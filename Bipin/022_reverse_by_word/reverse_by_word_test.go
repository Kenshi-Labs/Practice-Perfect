package main

import (
	"testing"
)

// Unit test for reverseByWord
func TestReverseByWord(t *testing.T) {
	// 1. Test with a sentence containing multiple words
	t.Run("Multiple words", func(t *testing.T) {
		input := "The sky is blue"
		expected := "blue is sky The"
		result := reverseByWord(input)
		if result != expected {
			t.Errorf("expected '%s', got '%s'", expected, result)
		}
	})

	// 2. Test with a single word
	t.Run("Single word", func(t *testing.T) {
		input := "Hello"
		expected := "Hello"
		result := reverseByWord(input)
		if result != expected {
			t.Errorf("expected '%s', got '%s'", expected, result)
		}
	})

	// 3. Test with a string containing leading/trailing spaces
	t.Run("Leading and trailing spaces", func(t *testing.T) {
		input := "  Hello world  "
		expected := "world Hello"
		result := reverseByWord(input)
		if result != expected {
			t.Errorf("expected '%s', got '%s'", expected, result)
		}
	})

	// 4. Test with an empty string
	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := ""
		result := reverseByWord(input)
		if result != expected {
			t.Errorf("expected '%s', got '%s'", expected, result)
		}
	})
}

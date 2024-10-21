package main

import (
	"testing"
)

// Test function for capitalizeFirstLetter
func TestCapitalizeFirstLetter(t *testing.T) {
	// Test case 1: Sentence with multiple words
	result := CapitalizeFirstLetter("the quick brown fox")
	expected := "The Quick Brown Fox"
	if result != expected {
		t.Errorf("Test Case 1 Failed: expected '%s', got '%s'", expected, result)
	}

	// Test case 2: Single word
	result = CapitalizeFirstLetter("hello")
	expected = "Hello"
	if result != expected {
		t.Errorf("Test Case 2 Failed: expected '%s', got '%s'", expected, result)
	}

	// Test case 3: String with numbers and special characters
	result = CapitalizeFirstLetter("123 hello! world")
	expected = "123 Hello! World"
	if result != expected {
		t.Errorf("Test Case 3 Failed: expected '%s', got '%s'", expected, result)
	}

	// Test case 4: Empty string
	result = CapitalizeFirstLetter("")
	expected = ""
	if result != expected {
		t.Errorf("Test Case 4 Failed: expected '%s', got '%s'", expected, result)
	}
}

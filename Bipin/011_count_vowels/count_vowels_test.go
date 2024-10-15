package main

import (
	"testing"
)

// Test function for countVowels
func TestCountVowels(t *testing.T) {
	// Test case 1: String with vowels
	result := CountVowels("Hello, World!")
	expected := 3
	if result != expected {
		t.Errorf("Test Case 1 Failed: expected %d, got %d", expected, result)
	}

	// Test case 2: String with no vowels
	result = CountVowels("xyz123")
	expected = 0
	if result != expected {
		t.Errorf("Test Case 2 Failed: expected %d, got %d", expected, result)
	}

	// Test case 3: Empty string
	result = CountVowels("")
	expected = 0
	if result != expected {
		t.Errorf("Test Case 3 Failed: expected %d, got %d", expected, result)
	}

	// Test case 4: String with only vowels
	result = CountVowels("aeiouAEIOU")
	expected = 10
	if result != expected {
		t.Errorf("Test Case 4 Failed: expected %d, got %d", expected, result)
	}
}

package main

import "testing"

// TestAreAnagrams contains unit tests for the AnagramChecker function.
func TestAnagramChecker(t *testing.T) {
	// 1. Test with valid anagrams
	if !AnagramChecker("Debit Card", "Bad Credit") {
		t.Errorf("Expected true, got false for 'Debit Card' and 'Bad Credit'")
	}

	// 2. Test with strings that are not anagrams
	if AnagramChecker("hello", "world") {
		t.Errorf("Expected false, got true for 'hello' and 'world'")
	}

	// 3. Test with strings of different lengths
	if AnagramChecker("abc", "abcd") {
		t.Errorf("Expected false, got true for 'abc' and 'abcd'")
	}

	// 4. Test with empty strings
	if !AnagramChecker("", "") {
		t.Errorf("Expected true, got false for two empty strings")
	}
}

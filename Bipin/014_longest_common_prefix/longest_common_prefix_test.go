package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// Test 1: Array with a common prefix
	result := LongestCommonPrefix([]string{"flower", "flow", "flight"})
	expected := "fl"
	if result != expected {
		t.Errorf("Test 1 failed. Expected %s, got %s", expected, result)
	}

	// Test 2: Array with no common prefix
	result = LongestCommonPrefix([]string{"dog", "racecar", "car"})
	expected = ""
	if result != expected {
		t.Errorf("Test 2 failed. Expected %s, got %s", expected, result)
	}

	// Test 3: Array containing an empty string
	result = LongestCommonPrefix([]string{"", "flow", "flight"})
	expected = ""
	if result != expected {
		t.Errorf("Test 3 failed. Expected %s, got %s", expected, result)
	}

	// Test 4: Array containing only one string
	result = LongestCommonPrefix([]string{"single"})
	expected = "single"
	if result != expected {
		t.Errorf("Test 4 failed. Expected %s, got %s", expected, result)
	}
}

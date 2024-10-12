package main

import "testing"

// TestRemoveVowels contains unit tests for the RemoveVowels function.
func TestRemoveVowels(t *testing.T) {
	// 1. Test with a string containing vowels
	input1 := "leetcode"
	expected1 := "ltcd"
	if result := RemoveVowels(input1); result != expected1 {
		t.Errorf("Expected %s, but got %s", expected1, result)
	}

	// 2. Test with a string containing no vowels
	input2 := "rhythm"
	expected2 := "rhythm"
	if result := RemoveVowels(input2); result != expected2 {
		t.Errorf("Expected %s, but got %s", expected2, result)
	}

	// 3. Test with an empty string
	input3 := ""
	expected3 := ""
	if result := RemoveVowels(input3); result != expected3 {
		t.Errorf("Expected %s, but got %s", expected3, result)
	}

	// 4. Test with a string containing only vowels
	input4 := "aeiouAEIOU"
	expected4 := ""
	if result := RemoveVowels(input4); result != expected4 {
		t.Errorf("Expected %s, but got %s", expected4, result)
	}
}

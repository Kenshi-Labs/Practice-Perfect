package main

import "testing"

func TestCountVowels(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{"Hello World", 3}, // contains vowels
		{"xyz123", 0},      // no vowels
		{"", 0},            // empty string
		{"aeiouAEIOU", 10},
	}

	for i, val := range testcases {
		if val.expected != CountVowel(val.input) {
			t.Errorf("Error in the test case no %d for the input of %s", i+1, val.input)
		}
	}
}

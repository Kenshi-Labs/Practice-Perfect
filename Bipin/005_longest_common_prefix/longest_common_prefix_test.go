package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// Define test cases using a struct
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"hello", "world", ""}, ""},
		{[]string{"only"}, "only"},
		{[]string{"", "", ""}, ""},
		{[]string{"", "test"}, ""},
		{[]string{"prefix", "pre", "preparation"}, "pre"},
		{[]string{"apple", "apricot", "ape"}, "ap"},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"mellow", "melon", "melancholy", "melt"}, "mel"},
		{[]string{"car", "cart", "card", "cartoon"}, "car"},
		{[]string{"testing", "testimony", "test", "tester"}, "test"},
		{[]string{"race", "racer", "racing", "racist"}, "rac"},
	}

	// Iterate over test cases
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := longestCommonPrefix(tc.input)
			if result != tc.expected {
				t.Errorf("Expected prefix %s, but got %s", tc.expected, result)
			}
		})
	}
}

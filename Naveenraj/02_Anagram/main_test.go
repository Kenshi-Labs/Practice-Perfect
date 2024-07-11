package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	testcases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "anagram", t: "nagaram", expected: true},
		{s: "ana", t: "nagaram", expected: false},
		{s: "seven", t: "eight", expected: false},
		{s: "", t: "", expected: true},
	}

	for i, val := range testcases {
		isAnagram := IsAnagram(val.s, val.t)
		if isAnagram != val.expected {
			t.Errorf("Error for the input %q", val.s)
		} else {
			fmt.Printf("Test Case %d Pass \n", i+1)
		}
	}
}

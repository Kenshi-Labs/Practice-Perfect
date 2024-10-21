package main

import (
	"fmt"
	"testing"
)

func TestCaptilizeFirstLetter(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"First leTTeR of EACH Word", "First Letter of Each Word"}, //containing multiple words
		{"single", "Single"},                     //single words
		{"123 hello! world", "123 Hello! World"}, //containing number and special character
		{"", ""},                                 //empty string
	}

	for i, val := range testcases {
		if val.expected != CaptilizeFirstLetter(val.input) {
			fmt.Printf("error in testcase numbet %d %v\n", i, t)
		}
	}
}

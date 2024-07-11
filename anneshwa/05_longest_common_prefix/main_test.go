package main

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	// Test with an array of strings with a common prefix
	input1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	output1 := LongestCommonPrefix(input1)
	assert.Equal(t, expected1, output1, "Test 1 Failed")

	// Test with an array of strings with no common prefix
	input2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	output2 := LongestCommonPrefix(input2)
	assert.Equal(t, expected2, output2, "Test 2 Failed")

	// Test with an array containing an empty string
	input3 := []string{"", "b", "c"}
	expected3 := ""
	output3 := LongestCommonPrefix(input3)
	assert.Equal(t, expected3, output3, "Test 3 Failed")

	// Test with an array containing only one string
	input4 := []string{"alone"}
	expected4 := "alone"
	output4 := LongestCommonPrefix(input4)
	assert.Equal(t, expected4, output4, "Test 4 Failed")
}
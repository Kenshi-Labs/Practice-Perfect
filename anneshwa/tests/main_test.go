package tests

import (
	"testing"

	functions "github.com/Kenshi-Labs/Practice-Perfect/pkg"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "olleh", functions.ReverseString("hello"))
	assert.Equal(t, "", functions.ReverseString(""))
	assert.Equal(t, "  dlrow olleh  ", functions.ReverseString("  hello world  "))
	assert.Equal(t, "racecar", functions.ReverseString("racecar"))
}

func TestIsAnagram(t *testing.T) {
	assert.True(t, functions.IsAnagram("anagram", "nagaram"))
	assert.True(t, functions.IsAnagram("world hello", "hello world"))
	assert.False(t, functions.IsAnagram("rat", "car"))
}

func TestTwoSum(t *testing.T) {
	assert.ElementsMatch(t, []int{0, 1}, functions.TwoSum([]int{2, 7, 11, 15}, 9))
	assert.ElementsMatch(t, []int{1, 2}, functions.TwoSum([]int{3, 2, 4}, 6))
}
func TestRemoveDuplicates(t *testing.T) {
	// Test with an array containing duplicates
	input1 := []int{1, 1, 2}
	expected1 := []int{1, 2}
	newLength1 := functions.RemoveDuplicates(input1)
	assert.Equal(t, 2, newLength1, "Test 1 Failed")
	assert.Equal(t, expected1, input1[:newLength1], "Test 1 Failed")

	// Test with an array with no duplicates
	input2 := []int{1, 2, 3}
	expected2 := []int{1, 2, 3}
	newLength2 := functions.RemoveDuplicates(input2)
	assert.Equal(t, 3, newLength2, "Test 2 Failed")
	assert.Equal(t, expected2, input2[:newLength2], "Test 2 Failed")

	// Test with an empty array
	input3 := []int{}
	newLength3 := functions.RemoveDuplicates(input3)
	assert.Equal(t, 0, newLength3, "Test 3 Failed")
	assert.Equal(t, []int{}, input3, "Test 3 Failed")

	// Test with an array containing all identical elements
	input4 := []int{2, 2, 2, 2}
	expected4 := []int{2}
	newLength4 := functions.RemoveDuplicates(input4)
	assert.Equal(t, 1, newLength4, "Test 4 Failed")
	assert.Equal(t, expected4, input4[:newLength4], "Test 4 Failed")
}

func TestLongestCommonPrefix(t *testing.T) {
	// Test with an array of strings with a common prefix
	input1 := []string{"flower", "flow", "flight"}
	expected1 := "fl"
	output1 := functions.LongestCommonPrefix(input1)
	assert.Equal(t, expected1, output1, "Test 1 Failed")

	// Test with an array of strings with no common prefix
	input2 := []string{"dog", "racecar", "car"}
	expected2 := ""
	output2 := functions.LongestCommonPrefix(input2)
	assert.Equal(t, expected2, output2, "Test 2 Failed")

	// Test with an array containing an empty string
	input3 := []string{"", "b", "c"}
	expected3 := ""
	output3 := functions.LongestCommonPrefix(input3)
	assert.Equal(t, expected3, output3, "Test 3 Failed")

	// Test with an array containing only one string
	input4 := []string{"alone"}
	expected4 := "alone"
	output4 := functions.LongestCommonPrefix(input4)
	assert.Equal(t, expected4, output4, "Test 4 Failed")
}

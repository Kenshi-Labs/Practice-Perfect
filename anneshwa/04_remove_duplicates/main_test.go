package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	// Test with an array containing duplicates
	input1 := []int{1, 1, 2}
	expected1 := []int{1, 2}
	newLength1 := RemoveDuplicates(input1)
	assert.Equal(t, 2, newLength1, "Test 1 Failed")
	assert.Equal(t, expected1, input1[:newLength1], "Test 1 Failed")

	// Test with an array with no duplicates
	input2 := []int{1, 2, 3}
	expected2 := []int{1, 2, 3}
	newLength2 :=  RemoveDuplicates(input2)
	assert.Equal(t, 3, newLength2, "Test 2 Failed")
	assert.Equal(t, expected2, input2[:newLength2], "Test 2 Failed")

	// Test with an empty array
	input3 := []int{}
	newLength3 := RemoveDuplicates(input3)
	assert.Equal(t, 0, newLength3, "Test 3 Failed")
	assert.Equal(t, []int{}, input3, "Test 3 Failed")

	// Test with an array containing all identical elements
	input4 := []int{2, 2, 2, 2}
	expected4 := []int{2}
	newLength4 := RemoveDuplicates(input4)
	assert.Equal(t, 1, newLength4, "Test 4 Failed")
	assert.Equal(t, expected4, input4[:newLength4], "Test 4 Failed")
}

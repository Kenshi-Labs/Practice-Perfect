package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	// Test case 1: Two non-empty sorted lists
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	expected1 := []int{1, 1, 2, 3, 4, 4}
	assert.Equal(t, expected1, mergeTwoLists(l1, l2), "Error in test case 1")

	// Test case 2: One empty list and one non-empty list
	l3 := []int{}
	l4 := []int{1, 2, 3}
	expected2 := []int{1, 2, 3}
	assert.Equal(t, expected2, mergeTwoLists(l3, l4), "Error in test case 2")

	// Test case 3: Two empty lists
	l5 := []int{}
	l6 := []int{}
	expected3 := []int{}
	assert.Equal(t, expected3, mergeTwoLists(l5, l6), "Error in test case 3")

	// Test case 4: Lists of different lengths
	l7 := []int{1, 2, 3}
	l8 := []int{4, 5}
	expected4 := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected4, mergeTwoLists(l7, l8), "Error in test case 4")
}

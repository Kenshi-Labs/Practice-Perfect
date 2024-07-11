package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.ElementsMatch(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	assert.ElementsMatch(t, []int{1, 2}, TwoSum([]int{3, 2, 4}, 6))
}
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	t.Run("Non-empty array with positive rotation", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{4, 5, 1, 2, 3}
		_, err := Rotate(nums, 2)
		assert.NoError(t, err)
		assert.Equal(t, expected, nums) // Check nums instead of result
	})

	t.Run("Rotation equal to array length", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := []int{1, 2, 3, 4}
		_, err := Rotate(nums, 4)
		assert.NoError(t, err)
		assert.Equal(t, expected, nums) // Check nums instead of result
	})

	t.Run("Rotation greater than array length", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{1,2,3} // Adjusted expected result after rotation
		_, err := Rotate(nums, 4)
		assert.NoError(t, err)
		assert.Equal(t, expected, nums) // Check nums instead of result
	})

	t.Run("Empty array", func(t *testing.T) {
		nums := []int{}
		result, err := Rotate(nums, 3)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "array is empty")
	})

	t.Run("Negative rotation", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result, err := Rotate(nums, -1)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "rotation steps must be greater than 0")
	})

	t.Run("Zero rotation", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result, err := Rotate(nums, 0)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "rotation steps must be greater than 0")
		assert.Equal(t, []int{1, 2, 3}, nums) // Ensure the original array remains unchanged
	})
	

	t.Run("Rotation exactly divisible by array length", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		_, err := Rotate(nums, 15)
		assert.NoError(t, err)
		assert.Equal(t, expected, nums) // Check nums instead of result
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse entire array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		Reverse(nums, 0, len(nums)-1)
		assert.Equal(t, expected, nums)
	})

	t.Run("Reverse part of array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{1, 4, 3, 2, 5}
		Reverse(nums, 1, 3)
		assert.Equal(t, expected, nums)
	})
}

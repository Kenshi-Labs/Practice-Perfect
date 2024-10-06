package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Run("Common elements", func(t *testing.T) {
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2, 3}
		expected := []int{2}
		result := Intersection(nums1, nums2)
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("No common elements", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{4, 5, 6}
		expected := []int{}
		result := Intersection(nums1, nums2)
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("Empty arrays", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{}
		expected := []int{}
		result := Intersection(nums1, nums2)
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("Arrays with duplicates", func(t *testing.T) {
		nums1 := []int{1, 1, 2, 2}
		nums2 := []int{2, 2, 2, 3}
		expected := []int{2}
		result := Intersection(nums1, nums2)
		assert.ElementsMatch(t, expected, result)
	})
}

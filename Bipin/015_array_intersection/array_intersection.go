package main

// Intersection function returns a new array containing the common elements from nums1 and nums2, without duplicates
func Intersection(nums1, nums2 []int) []int {
	// Use a map to track the presence of elements from nums1
	set := make(map[int]struct{})
	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	// Use another map to store the intersection elements (avoiding duplicates)
	intersectionSet := make(map[int]struct{})
	for _, num := range nums2 {
		if _, found := set[num]; found {
			intersectionSet[num] = struct{}{}
		}
	}

	// Convert the map to a slice
	result := []int{}
	for num := range intersectionSet {
		result = append(result, num)
	}

	return result
}

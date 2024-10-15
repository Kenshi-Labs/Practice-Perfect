package main

// RotateArray: Function to rotate an array to the right by k steps
func RotateArray(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// Reduce k to be within the array length
	k = k % n

	// Helper function to reverse a section of the array
	reverse := func(arr []int, start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	// Reverse the entire array
	// i.e. if input is 1, 2, 3, 4, 5 then after this function array will be 5, 4, 3, 2, 1
	reverse(nums, 0, n-1)

	// Reverse the first k elements
	// i.e. if after 1st reverse the array is 5, 4, 3, 2, 1 and k is 2 then after this function array will be 4, 5, 3, 2, 1
	reverse(nums, 0, k-1)

	// Reverse the rest of the array
	// i.e. if after 2nd reverse the array is 4, 5, 3, 2, 1 and k is 2 then after this function array will be 4, 5, 1, 2, 3
	reverse(nums, k, n-1)
}

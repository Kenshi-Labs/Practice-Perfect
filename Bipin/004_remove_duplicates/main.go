package main

import "fmt"

func main() {
	// Array with duplicates
	nums1 := []int{1, 1, 2}
	fmt.Println("Original nums1:", nums1)
	length1 := removeDuplicates(nums1)
	fmt.Println("After removing duplicates, nums1:", nums1[:length1])
	fmt.Println("New length:", length1)

	// Array with no duplicates
	nums2 := []int{1, 2, 3}
	fmt.Println("Original nums2:", nums2)
	length2 := removeDuplicates(nums2)
	fmt.Println("After removing duplicates, nums2:", nums2[:length2])
	fmt.Println("New length:", length2)

	// Empty array
	nums3 := []int{}
	fmt.Println("Original nums3:", nums3)
	length3 := removeDuplicates(nums3)
	fmt.Println("After removing duplicates, nums3:", nums3[:length3])
	fmt.Println("New length:", length3)

	// Array with all identical elements
	nums4 := []int{1, 1, 1, 1, 1}
	fmt.Println("Original nums4:", nums4)
	length4 := removeDuplicates(nums4)
	fmt.Println("After removing duplicates, nums4:", nums4[:length4])
	fmt.Println("New length:", length4)
}

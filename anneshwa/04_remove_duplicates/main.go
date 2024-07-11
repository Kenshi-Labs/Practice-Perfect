package main

import(
	"fmt"
)

func main(){
		// Testing RemoveDuplicates
		fmt.Println("Testing RemoveDuplicates:")
		arr1 := []int{1, 1, 2}
		newLength1 := RemoveDuplicates(arr1)
		fmt.Println(newLength1, arr1[:newLength1]) // Output: 2 [1 2]
	
		arr2 := []int{1, 2, 3}
		newLength2 := RemoveDuplicates(arr2)
		fmt.Println(newLength2, arr2[:newLength2]) // Output: 3 [1 2 3]
	
		arr3 := []int{}
		newLength3 := RemoveDuplicates(arr3)
		fmt.Println(newLength3, arr3[:newLength3]) // Output: 0 []
	
		arr4 := []int{2, 2, 2, 2}
		newLength4 := RemoveDuplicates(arr4)
		fmt.Println(newLength4, arr4[:newLength4]) // Output: 1 [2]
	}

	// PROBLEM 4: RemoveDuplicates function removes duplicates from a sorted array in-place
// and returns the new length.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}
	return uniqueIndex + 1
}

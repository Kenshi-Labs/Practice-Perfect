package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	newLength := RemoveDuplicates(nums)
	fmt.Printf("Output: %d, nums = %v\n", newLength, nums[:newLength])
}

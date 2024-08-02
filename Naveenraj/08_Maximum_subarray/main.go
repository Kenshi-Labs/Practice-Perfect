package main

import "fmt"

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	//finding all possible subarray
	for i := 0; i < len(nums); i++ {
		//i+1 because of j is exclusive
		for j := i + 1; j <= len(nums); j++ {
			sumOfArray := sumOfArray(nums[i:j])
			// if maximum number is lesser than sum we are assigning sum as max
			if max < sumOfArray {
				max = sumOfArray
			}
		}		
	} 
	return max
}

func sumOfArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	// iterating from [1:] because of 0th index added in sum already
	for _, val := range nums[1:] {
		sum += val
	}

	return sum
}
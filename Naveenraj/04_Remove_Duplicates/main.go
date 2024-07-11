package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(Remove_Duplicates(nums))
}

func Remove_Duplicates(nums []int) int {
	count := 1

	if len(nums) == 0 {
		return 0
	}

	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			nums[count] = nums[i]
			count++
		}

		prev = nums[i]
	}

	return count
}

func Remove_Duplicates_Original(nums []int) int {
	uniqueMap := make(map[int]bool)

	for _, val := range nums {
		if !uniqueMap[val] {
			uniqueMap[val] = true
		}
	}

	output := []int{}

	for k := range uniqueMap {
		output = append(output, k)
	}

	return len(output)
}

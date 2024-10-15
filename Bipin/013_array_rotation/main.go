package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	RotateArray(nums, k)
	fmt.Println(nums)
}

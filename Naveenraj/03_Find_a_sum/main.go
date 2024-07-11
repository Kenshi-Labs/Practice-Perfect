package main

import "fmt"

func main() {
	array := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(Find_a_sum(array, target))
}

func Find_a_sum(arr []int, targer int) []int {
	output := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == targer && i != j {
				output = []int{i, j}
				return output
			}
		}
	}
	return output
}

package main

import "fmt"

func main() {
	input := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(input))
}

func maxProfit(prices []int) int {
	max := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			temp := prices[j] - prices[i]
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

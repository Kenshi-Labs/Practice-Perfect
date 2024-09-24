package main

import "fmt"

func main() {
	// Optionally, manually test with a value
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Maximum profit from prices %v is: %d\n", prices, MaxProfit(prices))
}

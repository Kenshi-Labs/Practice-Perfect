package main

import "fmt"

// maxProfit calculates the maximum profit from the given stock prices.
func maxProfit(prices []int) int {
	// If there are no prices or only one price, no profit can be made.
	if len(prices) == 0 {
		return 0
	}

	// Initialize the minimum price to the first price and max profit to zero.
	minPrice := prices[0]
	maxProfit := 0

	// Iterate through the prices array to find the maximum profit.
	for _, price := range prices {
		// Update the minimum price if the current price is lower.
		if price < minPrice {
			minPrice = price
		}

		// Calculate the current profit by selling at the current price.
		profit := price - minPrice

		// Update the maximum profit if the current profit is greater.
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	// Return the maximum profit found.
	return maxProfit
}

func main() {
	// Example input
	prices := []int{7, 1, 5, 3, 6, 4}

	// Calculate the maximum profit
	profit := maxProfit(prices)

	// Output the result
	fmt.Printf("The maximum profit is: %d\n", profit)
}
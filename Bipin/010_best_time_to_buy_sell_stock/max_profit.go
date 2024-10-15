package main

// MaxProfit finds the maximum profit that can be achieved by buying and selling stock
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	// Loop through prices to find the maximum profit
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

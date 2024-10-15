package main

import "fmt"

func climbStairs(n int) int {
	// Edge cases
	if n == 0 {
		// There's 1 way to stay on the ground (do nothing)
		return 1
	} else if n == 1 {
		// There's only 1 way to reach the first step (take 1 step)
		return 1
	}

	// Initialize a slice to store number of ways to reach each step
	dp := make([]int, n+1)
	// There's 1 way to be on the ground (do nothing)
	dp[0] = 1
	// There's 1 way to reach the first step (take 1 step)
	dp[1] = 1

	// Fill the dp array according to the rules
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	n := 2
	fmt.Println(climbStairs(n)) // Output: 2
}

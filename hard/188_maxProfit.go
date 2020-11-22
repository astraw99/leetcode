package main

import "fmt"

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][][]int, n) // [day][k][0/1]
	for i, _ := range dp {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2) // hold or not
			dp[0][j][0] = 0
			dp[0][j][1] = -prices[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

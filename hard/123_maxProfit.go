package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}

// max 2 deals
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	K := 2
	dp := make([][][]int, n)
	for i, _ := range dp {
		dp[i] = make([][]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = make([]int, 2) // hold or not
			dp[0][j][0] = 0
			dp[0][j][1] = -prices[0]
		}
	}

	for i := 1; i < n; i++ {
		for k := K; k >= 1; k-- {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[n-1][K][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

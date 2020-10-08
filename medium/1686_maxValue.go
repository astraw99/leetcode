package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(grid))
}

func maxValue(grid [][]int) int {
	var res int
	if len(grid) == 0 {
		return res
	}

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

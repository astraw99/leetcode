package main

import (
	"fmt"
)

func main() {
	matrix := [][]byte{
		{'0', '0', '1', '1'},
		{'0', '1', '1', '0'},
		{'0', '1', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	var res int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
				if dp[i][j]*dp[i][j] > res {
					res = dp[i][j] * dp[i][j]
				}
			}
		}
	}

	return res
}

func min(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}

	return min
}

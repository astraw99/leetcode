package main

import "fmt"

func main() {
	m, n, k := 2, 3, 1
	fmt.Println(movingCount(m, n, k))
}

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, x, y, k int, dp [][]int) int {
	var res int

	if x < 0 || x > m-1 || y < 0 || y > n-1 || dp[x][y] == 1 || sumPos(x)+sumPos(y) > k {
		return 0
	}

	dp[x][y] = 1

	res += 1
	res += dfs(m, n, x+1, y, k, dp)
	res += dfs(m, n, x, y+1, k, dp)

	return res
}

func sumPos(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

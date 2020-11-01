package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChange2(coins, amount))
}

func coinChange(coins []int, amount int) interface{} {
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[j] = min(dp[j], dp[j-coins[i-1]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	var res = -1
	if len(coins) == 0 || amount == 0 {
		return res
	}

	// dp[i][j]:当考虑放入第i个物品时，占用空间j所能得到的最大的价值
	// 状态转移方程为：dp[i][j]=Max(dp[i-1][j],dp[i][j-vi]+wi)
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}

	for j := 0; j <= amount; j++ {
		dp[0][j] = amount + 1
	}

	dp[0][0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				// with coins[i]
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				// without coins[i]
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[len(coins)][amount] > amount {
		return res
	}

	return dp[len(coins)][amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

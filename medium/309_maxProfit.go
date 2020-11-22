package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit_(prices))
	fmt.Println(maxProfit(prices))
}

// with cool down (1 day)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dpI0, dpI1, dpIPre := 0, math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		temp := dpI0
		dpI0 = max(dpI0, dpI1+prices[i])
		dpI1 = max(dpI1, dpIPre-prices[i])
		dpIPre = temp
	}

	return dpI0
}

// with cool down (1 day)
func maxProfit_(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	memo := make(map[int]int)
	return dfs(prices, memo, 0)
}

func dfs(prices []int, memo map[int]int, start int) int {
	if start >= len(prices) {
		return 0
	}

	if v, ok := memo[start]; ok {
		return v
	}

	var res int
	minV := prices[start]

	for sell := start + 1; sell < len(prices); sell++ {
		minV = min(minV, prices[sell])
		res = max(res, dfs(prices, memo, sell+2)+prices[sell]-minV)
	}

	memo[start] = res

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

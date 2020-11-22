package main

import "fmt"

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var res int
	minV, maxV := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		if minV > prices[i] {
			minV = prices[i]
			maxV = prices[i]
		}

		if maxV < prices[i] {
			maxV = prices[i]
		}

		res = max(res, maxV-minV)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

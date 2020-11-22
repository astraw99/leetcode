package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
}

// with fee
func maxProfit(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dpI0, dpI1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		temp := dpI0
		dpI0 = max(dpI0, dpI1+prices[i]-fee)
		dpI1 = max(dpI1, temp-prices[i])
	}

	return dpI0
}

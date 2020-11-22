package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var res int
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - prices[i-1]; profit > 0 {
			res += profit
		}
	}

	return res
}

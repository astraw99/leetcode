package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var res int
	if len(prices) == 0 {
		return res
	}

	min, max := prices[0], prices[0]
	for i, _ := range prices {
		if i == 0 {
			continue
		}

		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min {
			min = prices[i]
			max = prices[i]
		}

		if res < max-min {
			res = max - min
		}
	}

	return res
}

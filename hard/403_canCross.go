package main

import (
	"fmt"
)

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(stones))
}

// 青蛙过河
func canCross(stones []int) bool {
	l := len(stones)

	dp := make(map[int]map[int]int, l)
	for i := 0; i < l; i++ {
		dp[stones[i]] = make(map[int]int)
	}

	dp[0][0] = 0

	for i := 0; i < l; i++ {
		for _, k := range dp[stones[i]] {
			for step := k - 1; step <= k+1; step++ {
				if step > 0 {
					if _, ok := dp[stones[i]+step]; ok {
						dp[stones[i]+step][i] = step
					}
				}
			}
		}
	}

	return len(dp[stones[l-1]]) > 0
}

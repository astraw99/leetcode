package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(10))
	fmt.Println(integerBreak(10))
}

// cuttingRope 动态规划
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		m := 0
		for j := i - 1; j > 0; j-- {
			m = max(dp[i-j]*j, m, (i-j)*j) // (i-j)*j: not cut anymore
		}
		dp[i] = m
		//fmt.Println(m)
	}
	return dp[n]
}

// integerBreak 贪心算法
func integerBreak(n int) int {
	if n <= 3 {
		return 1 * (n - 1)
	}
	res := 1
	if n%3 == 1 {
		res *= 4
		n -= 4
	}

	if n%3 == 2 {
		res *= 2
		n -= 2
	}
	for i := 0; i < n/3; i++ {
		res *= 3
	}
	return res
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}

	if b >= c && b >= a {
		return b
	}
	return c
}

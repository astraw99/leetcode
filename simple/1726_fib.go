package main

import "fmt"

func main() {
	fmt.Println(fib(10))
	fmt.Println(fib2(10))
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}

	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}

	return a
}

func fib2(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}

	return dp[n] % (1e9 + 7)
}

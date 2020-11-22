package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		// rob or not rob
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

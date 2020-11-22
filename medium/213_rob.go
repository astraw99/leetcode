package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

// optimize space after
func robRange(nums []int, start, end int) int {
	dpI, dpI1, dpI2 := 0, 0, 0
	for i := end; i >= start; i-- {
		// rob or not rob
		dpI = max(nums[i]+dpI2, dpI1)
		dpI2 = dpI1
		dpI1 = dpI
	}

	return dpI
}

// optimize space before
func robRange_(nums []int, start, end int) int {
	dp := make([]int, len(nums)+2)
	for i := end; i >= start; i-- {
		// rob or not rob
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}

	return dp[start]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(canPartition(nums))
	fmt.Println(canPartition_(nums))
}

func canPartition(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return false
	}

	sum, maxNum := 0, 0
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < l; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
			//fmt.Println(j, dp[j])
		}
	}

	return dp[target]
}

func canPartition_(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return false
	}

	var sum, maxNum int
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}

	// should be even number
	if sum&1 == 1 {
		return false
	}

	halfSum := sum / 2
	if maxNum > halfSum {
		return false
	}

	// 0 - 1 knapsack problem
	// dp[i][j] 表示从数组的 [0,i] 下标范围内选取若干个正整数（可以是 0 个），是否存在一种选取方案使得被选取的正整数的和等于 j
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, halfSum+1)
	}

	dp[0][nums[0]] = true
	for i := 1; i < l; i++ {
		dp[i][0] = true
		for j := 0; j <= halfSum; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[l-1][halfSum]
}

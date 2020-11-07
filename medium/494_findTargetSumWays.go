package main

import "fmt"

func main() {
	nums := []int{1, 0}
	S := 1
	fmt.Println(findTargetSumWays(nums, S))
	fmt.Println(findTargetSumWaysDp(nums, S))
	fmt.Println(findTargetSumWaysDfs(nums, S))
}

func findTargetSumWaysDp(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S {
		return 0
	}
	// 奇数不满足
	if (sum+S)&1 == 1 {
		return 0
	}

	target := (sum + S) >> 1

	// 有多少种方法装满 target 这个数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[target]
}

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}

	num := nums[0]
	nums = nums[1:]

	return findTargetSumWays(nums, S+num) + findTargetSumWays(nums, S-num)
}

// dfs right :)
func findTargetSumWaysDfs(nums []int, S int) int {
	var res int
	if len(nums) == 0 {
		return res
	}

	var cnt int
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if i == len(nums) {
			if sum == S {
				cnt++
			}
			return
		} else {
			dfs(i+1, sum+nums[i])
			dfs(i+1, sum-nums[i])
		}
	}

	dfs(0, 0)

	return cnt
}

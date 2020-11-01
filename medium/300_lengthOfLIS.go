package main

import "fmt"

func main() {
	nums := []int{3, 1, 2}
	fmt.Println(lengthOfLIS(nums))
	fmt.Println(lengthOfLIS_(nums))
}

/**
1. dp[i]表示以下标 i 为子序列结束时最长上升子序列的长度 初始时默认为1
2. 递推式为 dp[i] = max(dp[i],dp[j]+1) for j in 0 to i,其含义为:
   以nums[i]为结束的最长上升子序列是之前所有子序列中最长上升子序列长度+1的最大值
3. 最终的结果是dp数组中最大的一个值,其含义为
   一个数组的最长上升子序列的长度是其以任意一个元素为序列结束位置的长度最大值
*/
// 最长上升子序列（Longest Increasing Subsequence，LIS）
func lengthOfLIS(nums []int) int {
	var res int

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// self wrong!!
func lengthOfLIS_(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	min, max := nums[0], nums[0]
	m := make(map[int]int)
	m[max], m[min] = 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			max = nums[i]
			m[max] = i
		} else {
			if nums[i] < min {
				min = nums[i]
				m[min] = i
			}
		}
	}

	temp := 0
	for i := m[min]; i < m[max]; i++ {
		if nums[i+1] < nums[i] {
			temp++
		}
	}

	//fmt.Println(m, temp, min, max)

	return m[max] - m[min] - temp + 1
}

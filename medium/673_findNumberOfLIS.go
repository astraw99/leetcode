package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	fmt.Println(findNumberOfLIS(nums))
	fmt.Println(LIS(nums))
}

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	maxLen := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	fmt.Println(maxLen)

	res := 0
	for i := 0; i < len(cnt); i++ {
		if dp[i] == maxLen {
			res += cnt[i]
		}
	}

	return res
}

// 没有做出来具体的最小值 ??
func LIS(arr []int) []int {
	// write code here
	var res []int
	if len(arr) == 0 {
		return res
	}

	maxLen := 0
	dp := make([]int, len(arr))
	path := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		var temp []int
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				if dp[i] < dp[j]+1 {
					temp = append(temp, arr[j])
					dp[i] = dp[j] + 1
				}
				//dp[i] = max(dp[i], dp[j]+1)
			}
		}
		temp = append(temp, arr[i])

		maxLen = max(maxLen, dp[i])
		path[i] = temp
	}
	for i := 0; i < len(path); i++ {
		if dp[i] == maxLen {
			res = path[i]
		}
	}

	fmt.Println(path, maxLen)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

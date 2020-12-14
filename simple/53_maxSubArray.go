package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArray2(nums))
}

// greedy method
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum > 0 {
			curSum += nums[i]
		} else {
			curSum = nums[i]
		}

		maxSum = max(maxSum, curSum)
	}

	return maxSum
}

// dp method
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		maxSum = max(maxSum, nums[i])
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

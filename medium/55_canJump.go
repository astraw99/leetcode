package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
	fmt.Println(canJumpDP(nums))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	right := 0
	for i := 0; i < len(nums); i++ {
		// 如果跳不到位置i，直接返回false
		if right < i {
			return false
		}

		if i+nums[i] > right {
			right = i + nums[i]
		}

		if right >= len(nums)-1 {
			return true
		}
	}

	return false
}

func canJumpDP(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	var dp = make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		flag := false

		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}

		if flag {
			dp[i] = true
		}
	}

	return dp[len(nums)-1]
}

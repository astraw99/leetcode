package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 1}
	k := 3
	fmt.Println(subarraySumPre(nums, k))
	fmt.Println(subarraySum(nums, k))
}

// 前缀和 + 哈希表
func subarraySumPre(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, len(nums))
	m[0] = 1

	cnt, pre := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			cnt += v
		}

		m[pre]++
	}

	return cnt
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var cnt int
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}

	return cnt
}

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

// twoSum 一层循环写法
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
func twoSum(nums []int, target int) []int {
	var res []int

	length := len(nums)
	mapNums := make(map[int]int, length)

	for i := range nums {
		t := target - nums[i]
		if j, ok := mapNums[t]; ok {
			return []int{i, j}
		}

		mapNums[nums[i]] = i
	}

	return res
}

// twoSum2 两层循环写法
func twoSum2(nums []int, target int) []int {
	var res []int

	length := len(nums)
	start := 0
	end := length - 1

	for i := start; i < end; i++ {
		for j := end; j > i; j-- {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}

	return res
}

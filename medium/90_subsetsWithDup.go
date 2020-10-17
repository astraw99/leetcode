package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}

	// 先排序
	sort.Ints(nums)

	var dfs func(temp []int, cur int)
	dfs = func(temp []int, cur int) {
		res = append(res, append([]int{}, temp...))
		for i := cur; i < len(nums); i++ {
			// 同层去重
			if i > cur && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs([]int{}, 0)

	return res
}

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
	fmt.Println(subsetsMy(nums))
	fmt.Println(subsetsDFS(nums))
}

func subsetsMy(nums []int) [][]int {
	var res [][]int

	for mask := 0; mask < 1<<len(nums); mask++ {
		var set []int
		for i := 0; i < len(nums); i++ {
			if mask>>i&1 > 0 {
				set = append(set, nums[i])
				//fmt.Println(set)
			}
		}

		res = append(res, set)
	}

	return res
}

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
				//fmt.Println("in", set)
			}
		}
		//fmt.Println("out", set)

		ans = append(ans, append([]int(nil), set...))
	}

	return
}

func subsetsDFS(nums []int) (ans [][]int) {
	var set []int
	var dfs func(int)

	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)           // 取 nums[cur]
		set = set[:len(set)-1] // nums[cur] 进行回溯
		dfs(cur + 1)           // 不取 nums[cur]
	}

	dfs(0)

	return
}

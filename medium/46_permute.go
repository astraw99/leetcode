package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var res [][]int

	var backtrack = func(path []int) {}
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			// new []int
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			// repeat filter
			if inSlice(path, nums[i]) {
				continue
			}

			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})

	return res
}

func inSlice(nums []int, val int) bool {
	if len(nums) == 0 {
		return false
	}

	for _, v := range nums {
		if v == val {
			return true
		}
	}

	return false
}

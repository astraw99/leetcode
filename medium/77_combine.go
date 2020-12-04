package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

func combine(n int, k int) [][]int {
	var res [][]int
	if n <= 0 || k <= 0 {
		return res
	}

	var backtrack = func(path []int, start int) {}
	backtrack = func(path []int, start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := start; i <= n; i++ {
			// repeat filter
			if inSlice(path, i) {
				continue
			}

			path = append(path, i)
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{}, 1)

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

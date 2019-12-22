package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 5))
}

func search(nums []int, target int) int {
	res := -1
	for i := range nums {
		if nums[i] == target {
			res = i
			break
		}
	}

	return res
}

package main

import "fmt"

func main() {
	nums := []int{1}
	target := 1
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var first, last = -1, -1
	for i := 0; i < len(nums); i++ {
		if first < 0 && nums[i] == target {
			first = i
		}
		if last < 0 && nums[len(nums)-1-i] == target {
			last = len(nums) - 1 - i
		}

		if first >= 0 && last >= 0 {
			break
		}
	}

	return []int{first, last}
}

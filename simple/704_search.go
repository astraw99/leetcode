package main

import "fmt"

func main() {
	nums := []int{5}
	target := 5
	fmt.Println(search(nums, target))
}

// binarySearch 二分查找
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	res := -1
	for i := range nums {
		if nums[i] == target {
			res = i
			break
		}
	}

	return res
}

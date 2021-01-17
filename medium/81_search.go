package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}

// search 旋转有序数组(有重复值)
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1
	for left <= right {
		// delete duplicate
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}

		// left is sorted
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

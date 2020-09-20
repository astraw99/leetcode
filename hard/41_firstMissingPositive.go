package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(firstMissingPositive(nums))
}

// firstMissingPositive [LeetCode #41]
// 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数
func firstMissingPositive(nums []int) int {
	// 先排序
	sort.Ints(nums)
	length := len(nums)
	if length == 0 {
		return 1
	}
	if nums[0] > 0 && nums[0] != 1 {
		return 1
	}
	if nums[length-1] <= 0 {
		return 1
	}

	// 二分查找法找出 0 位置
	mid := length / 2
	fmt.Println(mid, nums)
	for mid > 0 {
		if nums[mid] > 0 {
			mid--
			if nums[mid] < 0 {
				return getMinPositive(nums[mid+1:], mid)
			}
		}
		if nums[mid] < 0 {
			mid++
			if nums[mid] > 0 {
				fmt.Println(mid)
				return getMinPositive(nums[mid:], mid)
			}
		}
		if nums[mid] == 0 {
			return getMinPositive(nums[mid+1:], mid)
		}
	}

	if mid == 0 {
		return getMinPositive(nums, mid)
	}

	return 1
}

// getMinPositive 取得最小正整数
func getMinPositive(nums []int, mid int) int {
	fmt.Println(nums)
	length := len(nums)
	mapNums := make(map[int]int, length)

	for i := 0; i < length; i++ {
		fmt.Println(mid, i)
		fmt.Println(mapNums, length, nums)

		if _, ok := mapNums[nums[i]]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			length = len(nums)
			continue
		}

		mapNums[nums[i]] = 1

		if nums[i] != i+1 {
			return i + 1
		}

		/*if nums[mid] > 0 && nums[mid+i] != i+1 {
			return i + 1
		}

		// -1, 1, 3, 4
		if nums[mid] < 0 {
			if nums[mid+i] > 0 && nums[mid+i] != i+1 {
				if nums[mid+i] > 1 {
					return 1
				}
				return i + 1
			}
		}

		if nums[mid] == 0 && nums[mid+i] != i {
			return i
		}*/
	}

	return nums[length-1] + 1
}

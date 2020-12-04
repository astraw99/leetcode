package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2, 1, 1, 0}
	sortColors(nums)
	sortColors2(nums)
}

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			if i != l {
				nums[i], nums[l] = nums[l], nums[i]
			}
			l++
		}
		if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			i--
			r--
		}

		//fmt.Println(nums, i)
	}

	fmt.Println(nums)
}

func sortColors2(nums []int) {
	/*sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})*/

	// 题目要求：一趟原地排序
	cntMap := make(map[int]int)
	for _, v := range nums {
		cntMap[v]++
	}

	lastIndex := 0
	for i := 0; i <= 2; i++ {
		if times, ok := cntMap[i]; ok {
			for repeat := 0; repeat < times; repeat++ {
				nums[lastIndex+repeat] = i
				//fmt.Println(nums, lastIndex)
			}
			lastIndex += times
		}
	}

	fmt.Println(nums)
}

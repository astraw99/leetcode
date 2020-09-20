package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().UnixNano()
	/*//nums := []int{4, 2, 1, 1, 2, 4, 5, 6, 6}
	nums := []int{8, 9, 9, 7, 7, 9, 9, 9}
	fmt.Println(singleNumber1(nums))
	fmt.Println(singleNumber2(nums))
	fmt.Println(majorityElement1(nums))
	fmt.Println(majorityElement2(nums))*/

	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix2(matrix, 22))

	endTime := time.Now().UnixNano()
	fmt.Printf("total time(ns): %d", endTime-startTime)
}

// searchMatrix1 二维矩阵搜索目标值
func searchMatrix1(matrix [][]int, target int) bool {
	ret := false
	for i, _ := range matrix {
		if ret {
			break
		}
		for _, inValue := range matrix[i] {
			if inValue == target {
				ret = true
				break
			}
		}
	}

	return ret
}

// searchMatrix2 二维矩阵搜索目标值
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	i, j := 0, col-1
	for i < row && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}

	return false
}

// majorityElement1 求众数1：众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素
func majorityElement1(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		equalNum := 1
		for j := 0; j < n; j++ {
			if nums[i] == nums[j] && i != j {
				equalNum++
			}

			if equalNum > n/2 {
				return nums[i]
			}
		}
	}

	return nums[0]
}

// majorityElement2 求众数2：众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素
func majorityElement2(nums []int) int {
	middle := len(nums) / 2
	equalMap := make(map[int]int, middle)

	for _, v := range nums {
		equalMap[v] += 1

		if equalMap[v] > middle {
			return v
		}
	}

	return nums[0]
}

// singleNumber1 只出现一次的数
func singleNumber1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i] == nums[j] && i != j {
				break
			}

			if j == (n - 1) {
				return nums[i]
			}
		}
	}

	return nums[n-1]
}

// singleNumber2 只出现一次的数
func singleNumber2(nums []int) int {
	// 零和任何数异或都等于任何数, 一个数异或两次就等于0
	// 本题中除一个之外每个元素都出现两次
	// 所以用循环异或所有数就等于 只出现一次的那个数
	ret := 0
	for _, v := range nums {
		ret ^= v
	}

	return ret
}

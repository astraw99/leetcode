package main

import (
	"fmt"
)

func main() {
	nums := []int{3, -1, 4}
	fmt.Println(maxProduct(nums))
	fmt.Println(maxProduct_(nums))
}

func maxProduct(nums []int) int {
	var res int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	minVal, maxVal, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tempMin := minVal
		minVal = min(min(nums[i], tempMin*nums[i]), maxVal*nums[i])
		maxVal = max(max(nums[i], tempMin*nums[i]), maxVal*nums[i])
		res = max(res, maxVal)
	}

	return res
}

// self wrong!!
func maxProduct_(nums []int) int {
	var res int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var negativeCnt int
	var negativePos []int
	var maxVal int
	res = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return max(max(maxProduct_(nums[:i]), 0), maxProduct_(nums[i+1:]))
		}

		if nums[i] < 0 {
			negativeCnt++
			negativePos = append(negativePos, i)
		}

		res *= nums[i]
		if maxVal < res {
			maxVal = res
		}
	}

	if negativeCnt%2 == 0 {
		return maxVal
	}

	if len(negativePos) > 0 {
		//maxPos := []int{0, negativePos[0], negativePos[len(negativePos) - 1], len(nums) - 1}
		maxVal = max(sliceProduct(nums[:negativePos[0]+1]), sliceProduct(nums[negativePos[0]+1:negativePos[len(negativePos)-1]+1]))
		maxVal = max(maxVal, sliceProduct(nums[negativePos[len(negativePos)-1]:]))
	}

	return maxVal
}

func sliceProduct(nums []int) int {
	var res = 1
	for i, _ := range nums {
		res *= nums[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

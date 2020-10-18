package main

import "fmt"

func main() {
	nums := []int{3, -1, 4}
	fmt.Println(maxProduct(nums))
}

// self wrong!!
func maxProduct(nums []int) int {
	var res int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var negativeCnt int
	var negativePos []int
	var max int
	res = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return maxFunc(maxFunc(maxProduct(nums[:i]), 0), maxProduct(nums[i+1:]))
		}

		if nums[i] < 0 {
			negativeCnt++
			negativePos = append(negativePos, i)
		}

		res *= nums[i]
		if max < res {
			max = res
		}
	}

	if negativeCnt%2 == 0 {
		return max
	}

	if len(negativePos) > 0 {
		//maxPos := []int{0, negativePos[0], negativePos[len(negativePos) - 1], len(nums) - 1}
		max = maxFunc(sliceProduct(nums[:negativePos[0]+1]), sliceProduct(nums[negativePos[0]+1:negativePos[len(negativePos)-1]+1]))
		max = maxFunc(max, sliceProduct(nums[negativePos[len(negativePos)-1]:]))
	}

	return max
}

func sliceProduct(nums []int) int {
	var res = 1
	for i, _ := range nums {
		res *= nums[i]
	}

	return res
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}

	return b
}

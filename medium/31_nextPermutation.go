package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 7, 1}
	nextPermutation(nums)
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	i, j, k := n-2, n-1, n-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for k >= 0 && nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	// reverse nums[j:end]
	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println(nums)
}

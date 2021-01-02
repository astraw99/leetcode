package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}

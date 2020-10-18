package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, 6, 4, 10, 9, 5}
	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var min, max = len(nums) - 1, 0
	var sortedNums = make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != sortedNums[i] {
			min = int(math.Min(float64(i), float64(min)))
			max = int(math.Max(float64(i), float64(max)))
		}
	}

	if max-min <= 0 {
		return 0
	}

	fmt.Println(min, max)

	return max - min + 1
}

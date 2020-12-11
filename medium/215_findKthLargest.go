package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, 0}
	k := 1
	fmt.Println(findKthLargest(nums, k))
	fmt.Println(findKthLargest2(nums, k))
}

func findKthLargest(nums []int, k int) int {
	// 单调递减队列 10 9 9 8
	if len(nums) == 0 || k <= 0 {
		return -1
	}

	var descQ []int
	for i := 0; i < len(nums); i++ {
		descQ = descendingQueue(descQ, nums[i], k)
	}

	if len(descQ) > 0 {
		return descQ[len(descQ)-1]
	}

	return -1
}

// wrong??
func descendingQueue(descQ []int, num, k int) []int {
	if len(descQ) == 0 {
		return []int{num}
	}

	if len(descQ) >= k {
		if descQ[len(descQ)-1] > num {
			return descQ
		}

		// 10 8 6  +7
		for i := len(descQ) - 2; i >= 0; i-- {
			if num >= descQ[0] {
				descQ = append([]int{num}, descQ[:len(descQ)-1]...)
				return descQ
			}

			if descQ[i] >= num {
				descQ = descQ[:len(descQ)-1]
				temp := append([]int{}, descQ[i+1:]...)
				descQ = append(descQ[:i+1], num)
				descQ = append(descQ, temp...)
				return descQ
			}
		}

		return descQ
	}

	for i := len(descQ) - 1; i >= 0; i-- {
		if num >= descQ[0] {
			return append([]int{num}, descQ...)
		}

		if num <= descQ[len(descQ)-1] {
			return append(descQ, num)
		}

		if descQ[i] >= num {
			temp := append([]int{}, descQ[i+1:]...)
			descQ = append(descQ[:i+1], num)
			descQ = append(descQ, temp...)
			return descQ
		}
	}

	return descQ
}

func findKthLargest2(nums []int, k int) int {
	if len(nums) <= 0 || k <= 0 {
		return -1
	}

	sort.Ints(nums)

	//fmt.Println(nums)

	return nums[len(nums)-k]
}

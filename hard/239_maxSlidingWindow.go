package main

import "fmt"

func main() {
	nums := []int{1, -1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}

	var queue []int
	for i := 0; i < len(nums); i++ {
		// 单调递减队列
		for i > 0 && len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, nums[i])
		// i-k >= 窗口最左边的数，则代表窗口已经满了
		if i >= k && nums[i-k] >= queue[0] {
			queue = queue[1:]
		}

		if i+1 >= k {
			res = append(res, queue[0])
		}
	}

	return res
}

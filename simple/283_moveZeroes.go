package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	moveZeroes2(nums)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	fmt.Println(nums)
}

func moveZeroes2(nums []int) {
	if len(nums) == 0 {
		return
	}

	slow, fast := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}

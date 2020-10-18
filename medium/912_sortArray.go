package main

import "fmt"

func main() {
	nums := []int{5, 2, 7, 3, 9, 2}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	//nums = quickSort(nums, 0, len(nums)-1)
	nums = bubbleSort(nums)

	return nums
}

func quickSort(nums []int, left, right int) []int {
	if len(nums) <= 1 {
		return nums
	}

	if left >= right {
		return nums
	}

	i, j := left, right
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}

	nums[i] = pivot
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)

	return nums
}

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		isSort := true
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isSort = false
			}
		}

		if isSort {
			break
		}
	}

	return nums
}

func swap(nums []int, i int, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp

	return nums
}

package main

import "fmt"

func main() {
	nums := []int{5, 2, 7, 3, 9, 2}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	nums = quickSort(nums, 0, len(nums)-1)
	nums = mergeSort(nums)
	nums = bubbleSort(nums)
	nums = bubbleSort2(nums)

	return nums
}

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 每一轮确定最大值
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(nums)
	}

	// 每一轮确定最小值
	/*for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		fmt.Println(nums)
	}*/

	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	nums = merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))

	return nums
}

func merge(nums1, nums2 []int) []int {
	var res []int
	n := len(nums1) + len(nums2)

	for i := 0; i < n; i++ {
		if len(nums1) == 0 {
			res = append(res, nums2...)
			return res
		}
		if len(nums2) == 0 {
			res = append(res, nums1...)
			return res
		}

		if nums1[0] <= nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}

	return res
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

func bubbleSort2(nums []int) []int {
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

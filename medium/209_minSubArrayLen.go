package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	sum := 7
	fmt.Println(minSubArrayLen(sum, nums))
	// todo 未通过，待继续
}

// minSubArrayLen [LeetCode #209]
// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。
// [2, 3, 1, 2, 4, 3] 4 + 3 >= 7 return = 2
// 如果不存在符合条件的连续子数组，返回 0。
func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	sum := 0
	minLength := 0
	step := length / 2 // 步长二分法

	for {
		for i := 0; i < length; i++ {
			if nums[i] >= s {
				return 1 // 直接满足条件的数
			}

			if i+step >= length {
				sum = sumSlice(nums[i:])
			} else {
				sum = sumSlice(nums[i : i+step])
			}

			if sum >= s {
				if minLength == 0 {
					minLength = step
				}
				if step < minLength {
					minLength = step
				}
				sum = 0
			}
		}

		if step == 2 {
			return minLength
		}

		if minLength == 0 {
			step += step / 2
		} else {
			step = step/2 + 1
		}

		if step >= length {
			break
		}
	}

	return minLength
}

// sumSlice 切片求和
func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

// 第一次写法
func minSubArrayLen2(s int, nums []int) int {
	length := len(nums)
	sum := 0
	addLength := 0
	minLength := 0

	for i := 0; i < length; i++ {
		if nums[i] >= s {
			return 1 // 直接满足条件的数
		}

		sum += nums[i]
		addLength++
		if sum >= s {
			if minLength == 0 {
				minLength = addLength
			}
			if addLength < minLength {
				minLength = addLength
			}
			addLength = 0
			sum = 0
		}

		fmt.Println(addLength, minLength)
	}

	sum = 0
	addLength = 0

	for i := length - 1; i >= 0; i-- {
		if nums[i] >= s {
			return 1 // 直接满足条件的数
		}

		sum += nums[i]
		addLength++
		if sum >= s {
			if minLength == 0 {
				minLength = addLength
			}
			if addLength < minLength {
				minLength = addLength
			}
			addLength = 0
			sum = 0
		}

		fmt.Println(addLength, minLength)
	}

	return minLength
}

package main

import "fmt"

func main() {
	nums := []int{24, 11, 49, 80, 63, 8, 61, 22, 73, 85}
	fmt.Println(partitionDisjoint(nums))
}

// partitionDisjoint [LeetCode #915]
// 给定一个数组 A，将其划分为两个不相交（没有公共元素）的连续子数组 left 和 right， 使得：
// left 中的每个元素都小于或等于 right 中的每个元素。
// left 和 right 都是非空的。
// left 要尽可能小。
// 在完成这样的分组后返回 left 的长度。可以保证存在这样的划分方法。
func partitionDisjoint(A []int) int {
	leftMax := A[0]
	max := A[0]
	rightIndex := 1
	length := len(A)

	for i := 1; i < length; i++ {
		if A[i] > max {
			max = A[i]
		}

		if A[i] < leftMax {
			rightIndex = i + 1
			leftMax = max
		}
	}

	fmt.Println(leftMax, max, rightIndex)

	return rightIndex
}

func partitionDisjoint2(A []int) int {
	var left []int
	var right []int

	leftMax := 0
	rightMin := 0
	rightMax := 0
	repeatTimes := 0
	repeatVal := 0

	for i, v := range A {
		if i == 0 {
			left = append(left, v)
			leftMax = v
			continue
		}

		if v < leftMax {
			left = append(left, v)
			if v < repeatVal {
				repeatTimes = 0
			}

			if v < rightMin {
				left = append(left, right...)
				leftMax = rightMax
				right = []int{}
			}
		}

		if v > leftMax {
			right = append(right, v)
			if v < rightMin {
				rightMin = v
			}
			if v > rightMax {
				rightMax = v
			}
			if rightMin == 0 {
				rightMin = v
				rightMax = v
			}
		}

		if v == leftMax {
			repeatTimes++
			repeatVal = v
			left = append(left, v)
		}
	}

	fmt.Println(left, right, repeatTimes)

	return len(left) - repeatTimes
}

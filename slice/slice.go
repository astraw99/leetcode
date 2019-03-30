package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge2(nums1, len(nums1)-len(nums2), nums2, len(nums2))
}

// 使用go-sort包排序
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Sort(sort.IntSlice(nums1))

	fmt.Println(nums1)
}

// 自己的原生写法
// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// 输出: [1,2,2,3,5,6]
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nonZero := 0
	for i, _ := range nums1 {
		if nums1[i] != 0 {
			nonZero = i
			break
		}
	}

	length := len(nums1)
	for i := 0; i < length; i++ {
		if nums1[i] == 0 {
			if length == 1 {
				nums1 = []int{}
				break
			}
			if i > nonZero {
				nums1 = append(nums1[:i], nums1[i+1:]...)
				i--
			}
			length = len(nums1)
		}
	}

	m = len(nums1)

	if m <= 0 || len(nums1) == 0 {
		nums1 = append(nums1, nums2...)
	} else {
		for i := 0; i < m; i++ {
			if len(nums1) == m+n {
				break
			}
			for index := 0; index < n; index++ {
				// corner case
				if nums2[0] >= nums1[m-1] && i == 0 && index == 0 {
					nums1 = append(nums1, nums2...)
					break
				}
				if nums2[n-1] <= nums1[0] && i == 0 && index == 0 {
					nums1 = append(nums2, nums1...)
					break
				}

				if nums2[index] > nums1[i] {
					if i == len(nums1)-1 {
						nums1 = append(nums1, nums2[index:]...)
						break
					}

					i++
					if index != n-1 {
						index--
					} else {
						nums1 = append(nums1, nums2[n-1])
						break
					}
					continue
				}
				if nums2[index] == nums1[i] {
					rear := append([]int{}, nums1[i+1:]...)
					nums1 = append(nums1[0:i+1], nums2[index])
					nums1 = append(nums1, rear...)

					i++
					continue
				}
				if nums2[index] < nums1[i] {
					if index == 0 && i == 0 {
						nums1 = append([]int{nums2[index]}, nums1...)
					} else {
						rear := append([]int{}, nums1[i:]...)
						nums1 = append(nums1[0:i], nums2[index])
						nums1 = append(nums1, rear...)
					}

					i++
					continue
				}
			}
		}
	}

	fmt.Println(nums1)
}

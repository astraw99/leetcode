package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}

// threeSum 滑尺中间靠方法
func threeSum(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	sort.Ints(nums)
	//fmt.Println(nums)

	mapRes := make(map[string]string) // 去重
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}

		j := i + 1
		k := length - 1
		for j < k {
			fmt.Println(nums[i], nums[j], nums[k], i, j, k)
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				mapKey := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
				if _, isExist := mapRes[mapKey]; !isExist {
					mapRes[mapKey] = "1"
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	fmt.Println(mapRes)

	return res
}

// threeSum 是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组
// 满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
func threeSum2(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	sort.Ints(nums)

	mapRes := make(map[string]string)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					// 1. map去重
					mapKey := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
					if _, isExist := mapRes[mapKey]; !isExist {
						mapRes[mapKey] = "1"
						res = append(res, []int{nums[i], nums[j], nums[k]})
					}

					// 2. reflect 去重耗时
					/*isExist := false
					for _, v := range res {
						if reflect.DeepEqual(v, []int{nums[i], nums[j], nums[k]}) {
							isExist = true
							break
						}
					}
					if !isExist {
						res = append(res, []int{nums[i], nums[j], nums[k]})
					}*/
				}
			}
		}
	}

	fmt.Println(nums, mapRes)

	return res
}

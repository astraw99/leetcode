package main

import "fmt"

func main() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumbers(nums))
}

func singleNumbers(nums []int) int {
	var res int

	mapExist := make(map[int]int)
	for i, _ := range nums {
		if v, ok := mapExist[nums[i]]; ok {
			if v == 2 {
				delete(mapExist, nums[i])
				continue
			}
		}
		mapExist[nums[i]] += 1
	}

	for i, _ := range mapExist {
		res = i
	}

	return res
}

package main

import "fmt"

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(nums))
	fmt.Println(singleNumbers2(nums))
}

func singleNumbers2(nums []int) []int {
	var res []int

	mapExist := make(map[int]int)
	for i, _ := range nums {
		if _, ok := mapExist[nums[i]]; ok {
			delete(mapExist, nums[i])
			continue
		}
		mapExist[nums[i]] = i
	}

	for i, _ := range mapExist {
		res = append(res, i)
	}

	return res
}

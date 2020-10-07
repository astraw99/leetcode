package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{3, 34, 5, 9, 30}
	fmt.Println(minNumber(nums))
}

func minNumber(nums []int) string {
	var res string
	if len(nums) == 0 {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) < strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})

	var numsStr []string
	for _, v := range nums {
		numsStr = append(numsStr, strconv.Itoa(v))
	}

	return strings.Join(numsStr, "")
}

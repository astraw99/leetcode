package main

import (
	"fmt"
	"sort"
)

func main() {
	shorter := 1
	longer := 2
	k := 3
	fmt.Println(divingBoard(shorter, longer, k))
}

// divingBoard [LeetCode #1666]
// todo 还没提交成功
// 你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。
// 你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。
// 返回的长度需要从小到大排列。
func divingBoard(shorter int, longer int, k int) []int {
	var res []int

	if k == 0 {
		return nil
	}

	for i := 1; i <= k; i++ {
		res = calSum(shorter, longer, i, res)
	}

	sort.Ints(res)

	return res
}

// calSum 计算长度和
func calSum(shorter, longer, k int, temp []int) []int {
	var res []int

	if k == 1 {
		if shorter == longer {
			temp = []int{shorter}
		} else {
			temp = []int{shorter, longer}
		}

		return temp
	}

	// 去重
	tempMap := make(map[int]int)
	for _, v := range temp {
		if _, ok := tempMap[v+shorter]; ok {
			continue
		} else {
			tempMap[v+shorter] = 1
		}

		if _, ok := tempMap[v+longer]; ok {
			continue
		} else {
			tempMap[v+longer] = 1
		}

		res = append(res, v+shorter, v+longer)
		/*if shorter == longer {
			res = append(res, v+shorter)
		} else {
		}*/
	}

	return res
}

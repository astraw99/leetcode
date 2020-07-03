package main

import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}

	k := 4
	fmt.Println(kthSmallest(matrix, k))
}

// kthSmallest [LeetCode #378] [参考了题解]
// 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
func kthSmallest(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, rows*cols)
	i := 0

	for _, row := range matrix {
		for _, val := range row {
			res[i] = val
			i++
		}
	}

	sort.Ints(res)
	fmt.Println(res)

	return res[k-1]
}

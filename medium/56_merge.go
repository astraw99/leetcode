package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 4},
		{0, 2},
		{3, 5},
	}
	fmt.Println(merge(intervals))
	fmt.Println(merge_(intervals))
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return res
	}

	// sort first
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		merged := res[len(res)-1]
		if intervals[i][0] > merged[1] {
			res = append(res, intervals[i])
			continue
		}

		if intervals[i][1] > merged[1] {
			merged[1] = intervals[i][1]
		}
	}

	return res
}

// self wrong!!
func merge_(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return res
	}

	if len(intervals) == 1 {
		return intervals
	}

	first := intervals[0]
	second := intervals[1]
	if second[0] <= first[1] && second[1] >= first[0] {
		merged := mergeInts(first, second)
		res = append(res, []int{merged[0], merged[1]})
	} else {
		res = append(res, first, second)
	}

	res = append(res, merge(intervals[2:])...)

	return res
}

func mergeInts(first, second []int) []int {
	min, max := first[0], first[0]
	for i := 0; i < len(first); i++ {
		if first[i] > max {
			max = first[i]
		}
		if first[i] < min {
			min = first[i]
		}
	}
	for i := 0; i < len(second); i++ {
		if second[i] > max {
			max = second[i]
		}
		if second[i] < min {
			min = second[i]
		}
	}

	return []int{min, max}
}

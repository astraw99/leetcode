/*
 * Copyright (C) 2023 https://github.com/astraw99. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License"); you may not use this
 * file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {0, 0}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		tmp := mergeTwo(intervals[i], intervals[i+1])
		if i == len(intervals)-2 {
			res = append(res, tmp...)
		} else {
			if len(tmp) == 1 {
				intervals[i+1] = tmp[0]
			} else {
				res = append(res, tmp[0])
			}
		}
	}

	return res
}

func mergeTwo(a, b []int) [][]int {
	if a[1] >= b[0] {
		return [][]int{{min(a[0], b[0]), max(a[1], b[1])}}
	}

	return [][]int{a, b}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

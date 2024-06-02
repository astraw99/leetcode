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
	//nums := []int{2, 7, 11, 15}
	//nums := []int{3, 2, 4}
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var res []int
	if len(nums) < 1 {
		return res
	}

	var numsMap = map[int]int{}
	for i, v := range nums {
		t := target - v
		if j, ok := numsMap[t]; ok {
			return []int{i, j}
		}
		numsMap[v] = i
	}

	return res
}

// failed!!
func twoSum2(nums []int, target int) []int {
	var res []int
	if len(nums) < 1 {
		return res
	}

	var numsMap = map[int]int{}
	for i, v := range nums {
		if _, ok := numsMap[v]; ok {
			continue
		}
		numsMap[v] = i
	}

	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			res = []int{numsMap[nums[i]], numsMap[nums[j]]}
			break
		}
		if sum < target {
			i++
		}
		if sum > target {
			j--
		}
	}

	return res
}

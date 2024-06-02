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
	nums := []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		tuples := twoSum(nums, i+1, 0-nums[i])
		for _, tuple := range tuples {
			res = append(res, append(tuple, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

func twoSum(nums []int, start, target int) [][]int {
	var res [][]int
	if len(nums) < 2 {
		return res
	}

	low, high := start, len(nums)-1
	for low < high {
		sum := nums[low] + nums[high]
		if sum < target {
			low++
			for low < high && nums[low] == nums[low-1] {
				low++
			}
		} else if sum > target {
			high--
			for low < high && nums[high] == nums[high+1] {
				high--
			}
		} else {
			res = append(res, []int{nums[low], nums[high]})
			low++
			high--
			for low < high && nums[low] == nums[low-1] {
				low++
			}
			for low < high && nums[high] == nums[high+1] {
				high--
			}
		}
	}

	return res
}

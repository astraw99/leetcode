/*
 * Copyright (C) https://github.com/astraw99. All rights reserved.
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

import "fmt"

func main() {
	//nums := []int{1, 2, 3}, k := 3
	//nums := []int{1, -1, 0}
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	fmt.Println(subarraySum(nums, k))
}

// 给你一个整数数组 nums 和一个整数 k，请你统计并返回 该数组中和为 k 的子数组的个数。
// 子数组是数组中元素的连续非空序列。
func subarraySum(nums []int, k int) int {
	var res int

	length := len(nums)
	if length < 1 {
		return res
	}

	for i := 0; i < length; i++ {
		var tmpSum int
		for j := i; j < length; j++ {
			tmpSum += nums[j]
			if k == tmpSum {
				res++
			}
		}
	}

	return res
}

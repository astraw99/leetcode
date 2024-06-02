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
	//nums := []int{100, 4, 200, 1, 3, 2}
	//nums := []int{1, 2, 3, 4, 100, 200}
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	var res, max = 1, 1

	if len(nums) < 1 {
		return 0
	}

	sort.Ints(nums)
	for i, _ := range nums {
		// skip duplicate num
		if i < (len(nums)-1) && nums[i+1] == nums[i] {
			continue
		}

		if i < (len(nums)-1) && nums[i+1]-nums[i] == 1 {
			res++
		} else {
			res = 1
		}
		max = maxInt(max, res)
	}

	return max
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

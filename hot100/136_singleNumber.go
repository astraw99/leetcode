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

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var res int
	if len(nums) == 0 {
		return res
	}

	var numsMap = map[int]int{}
	for i, _ := range nums {
		if _, ok := numsMap[nums[i]]; ok {
			numsMap[nums[i]] += 1
		} else {
			numsMap[nums[i]] = 1
		}
	}

	for i, v := range numsMap {
		if v == 1 {
			return i
		}
	}

	return res
}

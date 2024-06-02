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
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	backtrack := func(path []int) {}
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if inSlice(path, nums[i]) {
				continue
			}
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})

	return res
}

func inSlice(nums []int, v int) bool {
	for _, num := range nums {
		if num == v {
			return true
		}
	}
	return false
}

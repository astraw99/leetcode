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
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	right := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if right < i {
			return false
		}
		right = max(right, i+nums[i])
		if right >= n-1 {
			return true
		}
	}

	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

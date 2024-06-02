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
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var res int
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return nums[0]
	}

	res, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		if res < curSum {
			res = curSum
		}
	}

	return res
}

// O(N^2)
func maxSubArray2(nums []int) int {
	var res int
	length := len(nums)
	if length == 0 {
		return res
	}
	if length == 1 {
		return nums[0]
	}

	res = nums[0]
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if sum := sumArray(nums[i : j+1]); sum > res {
				res = sum
			}
		}
	}

	return res
}

func sumArray(nums []int) int {
	var res int
	for _, v := range nums {
		res += v
	}
	return res
}

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
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	//quickSort(nums, 0, len(nums)-1)
	bubbleSort(nums)
	return nums
}

func quickSort(nums []int, left, right int) {
	if len(nums) <= 1 {
		return
	}

	if left >= right {
		return
	}

	i, j := left, right
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}

	nums[i] = pivot
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func bubbleSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

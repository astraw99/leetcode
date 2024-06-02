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
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
	fmt.Println(dailyTemperatures2(temp))
	fmt.Println(dailyTemperatures3(temp))
}

// O(n) 单调栈
func dailyTemperatures(temperatures []int) []int {
	var res = make([]int, len(temperatures))
	if len(temperatures) == 0 {
		return res
	}

	n := len(temperatures)
	stack := make([]int, len(temperatures))

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			res[lastIndex] = i - lastIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

// O(n^2)
func dailyTemperatures2(temperatures []int) []int {
	var res = make([]int, len(temperatures))
	if len(temperatures) == 0 {
		return res
	}

	n := len(temperatures)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}

// O(n^2)
func dailyTemperatures3(temperatures []int) []int {
	var res []int
	if len(temperatures) == 0 {
		return res
	}

	n := len(temperatures)

	for i := 0; i < n; i++ {
		left, right := i, i+1
		tmp := 0
		for right < n {
			tmp++
			if temperatures[right] > temperatures[left] {
				break
			}
			right++
		}
		if right == n {
			tmp = 0
		}
		res = append(res, tmp)
	}

	return res
}

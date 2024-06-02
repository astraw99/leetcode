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
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	var res string
	if s == "" {
		return res
	}

	length := len(s)
	var left, right, maxStart, maxLen int

	for i := 0; i < length; i++ {
		tmpLen := 1
		left = i - 1
		right = i + 1
		for left >= 0 && s[i] == s[left] {
			left--
			tmpLen++
		}
		for right < length && s[i] == s[right] {
			right++
			tmpLen++
		}
		for left >= 0 && right < length && s[left] == s[right] {
			left--
			right++
			tmpLen += 2
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
			maxStart = left
		}
	}

	return s[maxStart+1 : maxStart+maxLen+1]
}

// self wrong!!
func longestPalindrome2(s string) string {
	var res string
	if s == "" {
		return res
	}

	for i := 0; i < len(s); i++ {
		tmpLen := 1
		low, high := i-1, i+1
		for low >= 0 && high < len(s) {
			if s[low] == s[high] {
				tmpLen += 2
				low--
				high++
			} else {
				break
			}
		}
		if tmpLen > len(res) {
			res = s[low+1 : high]
		}
	}

	return res
}

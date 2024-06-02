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
	"strings"
)

func main() {
	s1 := "hofubmnylkra"
	s2 := "pqhgxgdofcvmr"
	fmt.Println(longestCommonSubsequence(s1, s2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var res int

	if len(text1) == 0 || len(text2) == 0 {
		return res
	}

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

// error!!
func longestCommonSubsequence2(text1 string, text2 string) int {
	var res int

	if len(text1) == 0 || len(text2) == 0 {
		return res
	}

	if len(text2) > len(text1) {
		text1, text2 = text2, text1
	}

	for j := 0; j < len(text2); j++ {
		tmpLen := 0
		for start, searchIndex := 0, j; start < len(text1) && searchIndex < len(text2); {
			if hitIndex := strings.Index(text1, string(text2[searchIndex])); hitIndex >= 0 && hitIndex >= start {
				tmpLen++
				searchIndex++
				start = hitIndex
			} else {
				break
			}
		}
		res = max(res, tmpLen)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	var res int

	l := len(grid)
	if l == 0 {
		return res
	}

	var dp [][]int
	dp = make([][]int, l)
	for i, row := range grid {
		dp[i] = make([]int, len(row))
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < l; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[l-1][len(grid[l-1])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

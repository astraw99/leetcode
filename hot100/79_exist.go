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
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	var backtrack func(i, j, start int) bool

	backtrack = func(i, j, start int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '*' || board[i][j] != word[start] {
			return false
		}
		if start == len(word)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = '*'
		defer func() {
			board[i][j] = tmp
		}()
		start++
		return backtrack(i-1, j, start) ||
			backtrack(i+1, j, start) ||
			backtrack(i, j-1, start) ||
			backtrack(i, j+1, start)
	}

	for i, row := range board {
		for j := 0; j < len(row); j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

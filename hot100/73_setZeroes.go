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
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	var zeroMapRow, zeroMapCol = map[int]struct{}{}, map[int]struct{}{}

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				if _, ok := zeroMapRow[i]; !ok {
					zeroMapRow[i] = struct{}{}
				}
				if _, ok := zeroMapCol[j]; !ok {
					zeroMapCol[j] = struct{}{}
				}
			}
		}
	}

	for i, row := range matrix {
		for j := range row {
			_, ok1 := zeroMapRow[i]
			_, ok2 := zeroMapCol[j]
			if ok1 || ok2 {
				matrix[i][j] = 0
			}
		}
	}

	//fmt.Println("==========", zeroMapRow, zeroMapCol)

	/*for i, _ := range zeroMapRow {
		for rIdx, row := range matrix {
			if rIdx == i {
				for idx, _ := range row {
					row[idx] = 0
				}
				matrix[i] = row
				break
			}
		}
	}

	//fmt.Println("11111", matrix)

	for c, _ := range zeroMapCol {
		for rIdx, row := range matrix {
			for cIdx, _ := range row {
				if cIdx == c {
					matrix[rIdx][c] = 0
				}
			}
		}
	}*/
}

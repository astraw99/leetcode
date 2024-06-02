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
	"strconv"
	"strings"
)

func main() {
	// 输入：s = "3[a]2[bc]"
	//输出："aaabcbc"
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	var (
		numStack []int
		strStack []string
		res      string
		num      int
	)
	if len(s) == 0 {
		return res
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n, _ := strconv.Atoi(string(s[i]))
			num = num*10 + n
		} else if s[i] == '[' {
			strStack = append(strStack, res)
			res = ""
			numStack = append(numStack, num)
			num = 0
		} else if s[i] == ']' {
			cnt := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = str + strings.Repeat(res, cnt)
		} else {
			res += string(s[i])
		}
	}

	return res
}

// self wrong!!
func decodeString2(s string) string {
	var res string
	if len(s) == 0 {
		return res
	}

	repeat := 0
	repeatStr := ""
	var stack []byte
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			repeat, _ = strconv.Atoi(string(s[i]))
		} else if s[i] == '[' {
			stack = append(stack, s[i])
		} else if s[i] == ']' && len(stack) >= 1 {
			lastIndex := strings.LastIndex(string(stack), string('['))
			if lastIndex == 0 {
				repeatStr = string(stack[lastIndex:])
				for j := 0; j < repeat; j++ {
					res += repeatStr
				}
				repeat = 0
				stack = []byte{}
			} else {
				repeatStr = string(stack[lastIndex:])
				tmp := ""
				for j := 0; j < repeat; j++ {
					tmp += repeatStr
				}
				repeat = 0
				stack = append(stack[:lastIndex-1], tmp...)
			}

		} else {
			stack = append(stack, s[i])
		}
	}

	return res
}

func isDigit(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

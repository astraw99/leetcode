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
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var res int
	if len(prices) < 2 {
		return res
	}

	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}

	return res
}

// low performance!!
func maxProfit2(prices []int) int {
	var res int
	if len(prices) < 2 {
		return res
	}

	pricesLen := len(prices)
	for i := 0; i < pricesLen; i++ {
		for j := i + 1; j < pricesLen; j++ {
			if profit := prices[j] - prices[i]; profit > 0 {
				if profit > res {
					res = profit
				}
			}
		}
	}

	return res
}

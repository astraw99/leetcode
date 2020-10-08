package main

import "fmt"

func main() {
	n := 1352
	fmt.Println(nthUglyNumber(n))
	fmt.Println(nthUglyNumber_(n))
}

// dp: O(n)
func nthUglyNumber(n int) int {
	var res int
	if n <= 0 {
		return res
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	n2, n3, n5 := 1, 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = min(min(dp[n2]*2, dp[n3]*3), dp[n5]*5)
		if dp[i] == dp[n2]*2 {
			n2++
		}
		if dp[i] == dp[n3]*3 {
			n3++
		}
		if dp[i] == dp[n5]*5 {
			n5++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// self right. But O(n^2)
func nthUglyNumber_(n int) int {
	var res int
	if n <= 0 {
		return res
	}
	if n == 1 {
		return 1
	}

	mods := []int{2, 3, 5}
	i := 2
	num := 2
	temp := num
	for i <= n {
		temp = num

		for _, mod := range mods {
			flag := false
			for num%mod == 0 {
				if num/mod == 1 {
					i++
					flag = true
					break
				}
				num /= mod
				continue
			}

			if flag {
				break
			}
		}

		num = temp + 1
	}

	return temp
}

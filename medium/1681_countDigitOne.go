package main

import "fmt"

func main() {
	n := 2304
	fmt.Println(countDigitOne(n))
}

// 2304 high: 23, cur: 0, low: 4
// 2314 high: 23, cur: 1, low: 4
// 2324 high: 23, cur: 2, low: 4
func countDigitOne(n int) int {
	var res int
	var digit = 1
	var low = 0
	var high = n / 10
	var cur = n % 10

	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}

		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}

	return res
}

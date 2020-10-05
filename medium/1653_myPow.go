package main

import (
	"fmt"
	"math"
)

func main() {
	x := -13.62608
	n := 7
	fmt.Println(myPow(x, n))
	fmt.Println(myPowIteration(x, n))
	fmt.Println(myPowSelf(x, n))
}

// base = x, exponent = n
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.0
	for n >= 1 {
		if n&1 == 1 {
			res *= x
			n--
		} else {
			res *= x * x
			n = n >> 1
		}
	}

	return res
}

// base = x, exponent = n
func myPowIteration(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	temp := myPowIteration(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}

	return temp * temp * x
}

// base = x, exponent = n
func myPowSelf(x float64, n int) float64 {
	var res float64 = 1
	if x == 1 || n == 0 {
		return res
	}
	if x == 0 {
		return 0
	}

	// 判断 x 符号
	negativeX := false
	if x < 0 {
		x = -x
		negativeX = true
	}

	// 判断 n 符号
	negativeN := false
	if n < 0 {
		n = -n
		negativeN = true
	}

	if n >= math.MaxInt32 {
		if 0 < x && x < 1 {
			return 0
		}
		if x == 1 && negativeX {
			if n%2 == 0 {
				return res
			}
			return -res
		}
		if x > 1 {
			if negativeN {
				return 0
			}

			// x 为负数
			if negativeX {
				if n%2 == 0 {
					return math.MaxFloat64
				}
				return 0
			}
			return math.MaxFloat64
		}
	}

	for i := 0; i < n; i++ {
		res *= x
	}

	if negativeN {
		res = 1 / res
	}

	if negativeX {
		if n%2 != 0 {
			res = -res
		}
	}

	return res
}

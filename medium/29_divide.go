package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-56789, -12))
}

// divide 给定两个整数，被除数 dividend 和除数 divisor
// 将两数相除，要求不使用乘法、除法和 mod 运算符
func divide(dividend int, divisor int) int {
	sign := (dividend ^ divisor) < 0 // 异或取正负号

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	res := 0
	for i := 31; i >= 0; i-- {
		if dividend>>i >= divisor {
			res += 1 << i
			//fmt.Println(dividend>>i, divisor<<i)
			dividend -= divisor << i // divisor << i == divisor * (1 << i)
		}
	}

	if sign {
		return -res
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}

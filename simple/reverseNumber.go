package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}

// reverse 整数反转 [LeetCode #7]
func reverse(x int) int {
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	resStr := ""
	str := strconv.Itoa(x)
	length := len(str)

	for i := length - 1; i >= 0; i-- {
		resStr += string(str[i])
	}

	res, _ := strconv.Atoi(resStr)
	if res > math.MaxInt32 {
		res = 0
	}

	if sign {
		res = -res
	}

	return res
}

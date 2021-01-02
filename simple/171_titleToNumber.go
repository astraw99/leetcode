package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("AAA"))
	fmt.Println(titleToNumber2("AAA"))
}

func titleToNumber(s string) int {
	var res int
	if len(s) == 0 {
		return res
	}

	for i := 0; i < len(s); i++ {
		temp := int(s[i]-'A') + 1
		res = res*26 + temp
	}

	return res
}

func titleToNumber2(s string) int {
	var res int
	if len(s) == 0 {
		return res
	}

	for i := 0; i < len(s); i++ {
		temp := int(s[i]-'A') + 1
		if i < len(s)-1 {
			x := float64(26)
			y := float64(len(s) - 1 - i)
			temp *= int(math.Pow(x, y))
		}
		res += temp
	}

	return res
}

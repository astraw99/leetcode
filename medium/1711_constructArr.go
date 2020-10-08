package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(a))
}

func constructArr(a []int) []int {
	l := len(a)
	var res = make([]int, l)
	if len(a) == 0 {
		return res
	}

	left := make([]int, l)
	right := make([]int, l)
	left[0] = 1
	right[l-1] = 1

	for i := 1; i < l; i++ {
		left[i] = left[i-1] * a[i-1]
	}

	for j := l - 2; j >= 0; j-- {
		right[j] = right[j+1] * a[j+1]
	}

	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

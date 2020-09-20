package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "123"
	s2 := "456"
	fmt.Println(multiply(s1, s2))
}

func multiply(num1 string, num2 string) string {
	n1 := str2intArr(num1) // 倒序[6, 5, 4] [3, 2, 1]
	n2 := str2intArr(num2)
	if n1 == nil || n2 == nil {
		return ""
	}

	r := make([]int64, len(n1)+len(n2))

	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			r[i+j+1] += (r[i+j] + n1[i]*n2[j]) / 10
			r[i+j] = (r[i+j] + n1[i]*n2[j]) % 10
		}
	}

	var s string

	j := len(r) - 1
	for j > 0 && r[j] == 0 {
		j--
	}
	for i := 0; i <= j; i++ {
		s = fmt.Sprintf("%d", r[i]) + s
	}
	return s
}

func str2intArr(num string) []int64 {
	if len(num) == 0 {
		return nil
	}
	var r []int64
	for i := 0; i < len(num); i++ {
		v, err := strconv.ParseInt(string([]byte{num[i]}), 10, 64)
		if err != nil {
			return nil
		}
		r = append([]int64{v}, r...)
	}
	return r
}

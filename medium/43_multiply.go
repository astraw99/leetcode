package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "123"
	s2 := "456"
	fmt.Println(multiply(s1, s2))
	fmt.Println(strMultiply(s1, s2))
}

func strMultiply(num1, num2 string) string {
	var res string

	// num1[i] * num[j] 值必定在resArr[i+j] resArr[i+j+1]上，i+j+1 存储低位
	resArr := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			temp := int(num1[i]-'0')*int(num2[j]-'0') + resArr[i+j+1]
			resArr[i+j+1] = temp % 10
			resArr[i+j] += temp / 10
		}
	}

	// 去除前导 0
	i := 0
	for ; i < len(resArr); i++ {
		if resArr[i] != 0 {
			break
		}
	}
	for ; i < len(resArr); i++ {
		res += string(resArr[i] + '0') // 转换成 ASCII
	}
	if res == "" {
		return "0"
	}

	return res
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

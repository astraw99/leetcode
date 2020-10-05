package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 19
	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) int {
	// 写法一
	var start = 1
	var digit = 1
	var flag = 9

	for n-flag*digit > 0 {
		n -= flag * digit
		flag *= 10
		start *= 10
		digit++
	}

	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')

	// 写法二
	/*digits := 1
	flag := 9
	// 求n所在的数为几位数
	for n-flag*digits > 0 {
		n = n - flag*digits
		flag = flag * 10
		digits++
	}
	// 若n为个位数，返回n即可
	if digits == 1 {
		return n
	}
	number := 1
	for k := 1; k < digits; k++ {
		number = number * 10
	}
	// 求n对应的数
	number = number + (n-1)/digits
	idx := (n - 1) % digits
	// 将n对应的数转换为字符串类型，取第idx位并转换为整数
	strnums := strconv.Itoa(number)
	res, _ := strconv.Atoi(string(strnums[idx]))

	return res*/
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num1 := "999"
	num2 := "999"

	fmt.Println(multiply(num1, num2))
}

// multiply 字符串相乘
// num1 和 num2 的长度小于 110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
func multiply(num1 string, num2 string) string {
	res := ""

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)
	res = strings.Repeat("0", len1+len2) // 00000

	num1Bytes := []byte(num1)
	num2Bytes := []byte(num2)
	resBytes := []byte(res)

	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			num1Int, _ := strconv.Atoi(string(num1Bytes[i]))
			num2Int, _ := strconv.Atoi(string(num2Bytes[j]))
			curTemp, _ := strconv.Atoi(string(resBytes[i+j+1]))
			temp := curTemp + num1Int*num2Int
			resBytes[i+j+1] = []byte(strconv.Itoa(temp % 10))[0]
			carry, _ := strconv.Atoi(string(resBytes[i+j]))
			carry += temp / 10
			carryBytes := []byte(strconv.Itoa(carry))
			/*carryBytesLen := len(carryBytes)
			for m := 0; m < carryBytesLen; m++ {
				carryBytesLess, _ := strconv.Atoi(string(resBytes[i+j-m]))
				carryBytesLessTemp, _ := strconv.Atoi(string(carryBytes[0]))

				resBytes[i+j-m] = []byte(strconv.Itoa(carryBytesLess + carryBytesLessTemp))[0]
			}*/
			if carry >= 10 {
				if i+j >= 1 {
					carryBytesLess, _ := strconv.Atoi(string(resBytes[i+j-1]))
					carryBytesLessTemp, _ := strconv.Atoi(string(carryBytes[0]))
					resBytes[i+j-1] = []byte(strconv.Itoa(carryBytesLess + carryBytesLessTemp))[0] // 保存进位 carry
				}

				if i+j >= 2 {
					carryBytesLess2, _ := strconv.Atoi(string(resBytes[i+j-1-1]))
					carryBytesLessTemp2, _ := strconv.Atoi(string(carryBytes[0]))
					resBytes[i+j-1-1] = []byte(strconv.Itoa(carryBytesLess2 + carryBytesLessTemp2))[0] // 保存进位 carry
				}

				resBytes[i+j] = carryBytes[1] // 保存进位 carry
			} else {
				resBytes[i+j] = carryBytes[0]
			}

			fmt.Println(string(resBytes))
		}
	}

	res = string(resBytes)
	res = strings.TrimLeft(res, "0")

	return res
}

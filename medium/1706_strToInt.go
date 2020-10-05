package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "20000000000000000000"
	fmt.Println(strToInt(s))
}

func strToInt(str string) int {
	var res int
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return res
	}

	// 符号处理
	sign := true
	if !isDigit(str[0]) && str[0] != '+' && str[0] != '-' {
		return res
	}

	if str[0] == '+' || str[0] == '-' {
		if str[0] == '-' {
			sign = false
		}
		str = str[1:]
	}

	// 按位处理数字
	var temp []byte
	for i := 0; i < len(str); i++ {
		if isDigit(str[i]) {
			temp = append(temp, str[i])
		} else {
			break
		}
	}

	if len(temp) == 0 {
		return res
	}

	//fmt.Println(temp)

	isFirstZero := true
	for i := 0; i < len(temp); i++ {
		// '0' 对应 ascii = 48
		if isFirstZero && int(temp[i]-48) == 0 {
			continue
		} else {
			isFirstZero = false
		}

		if math.Pow10(len(temp)-1-i) > math.MaxInt32 {
			if sign {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		res += int(temp[i]-48) * int(math.Pow10(len(temp)-1-i))

		// res = res * 10 + int(temp[i]-'0') // 更简洁的写法
	}

	if !sign {
		res = -res
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	if res < math.MinInt32 {
		res = math.MinInt32
	}

	return res
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

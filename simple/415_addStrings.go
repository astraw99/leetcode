package main

import (
	"fmt"
)

func main() {
	num1 := "1"
	num2 := "99"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	var res string

	var carry int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(num1[i] - '0')
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
		}

		res = fmt.Sprint(carry%10) + res
		//res = strconv.Itoa(carry%10) + res
		carry /= 10
	}

	return res
}

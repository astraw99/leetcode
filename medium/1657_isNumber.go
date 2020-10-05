package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "3.e5"
	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	var res bool
	s = strings.Trim(s, " ")

	res, _ = regexp.MatchString("^(([+-]?[0-9]+(\\.[0-9]*)?)|([+-]?\\.?[0-9]+))([eE][+-]?[0-9]+)?$", s)

	return res
}

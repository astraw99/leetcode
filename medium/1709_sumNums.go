package main

import "fmt"

func main() {
	n := 10
	fmt.Println(sumNums(n))
}

func sumNums(n int) int {
	var res int

	var sum func(n int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}

	sum(n)

	return res
}

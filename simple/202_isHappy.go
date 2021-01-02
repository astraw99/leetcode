package main

import "fmt"

func main() {
	fmt.Println(isHappy(13))
}

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	var m = make(map[int]struct{})
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}

		m[n] = struct{}{}
		n = squareSum(n)
	}

	return true
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}

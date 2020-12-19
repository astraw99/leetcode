package main

import "fmt"

func main() {
	num := 10
	fmt.Println(hammingWeight(uint32(num)))
	fmt.Println(hammingWeight2(uint32(num)))
}

func hammingWeight(num uint32) int {
	var res int
	if num == 0 {
		return res
	}

	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}

	return res
}

func hammingWeight2(num uint32) int {
	var res int
	if num == 0 {
		return res
	}

	for num > 0 {
		num &= num - 1
		res++
	}

	return res
}

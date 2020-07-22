package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums))
}

func decompressRLElist(nums []int) []int {
	var res []int
	length := len(nums)

	for i := 0; i < length/2; i++ {
		freq, val := nums[2*i], nums[2*i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}

	return res
}

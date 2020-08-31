package main

import "fmt"

func main() {
	nums := []int{1, 2, -3, 5, -6, 4, 10, -7, 20}
	fmt.Println(getMaxLen(nums))
}

func getMaxLen(nums []int) int {
	var negativeIndex []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			l := getMaxLen(nums[:i])
			if l > len(nums[i+1:]) {
				return l
			}

			r := getMaxLen(nums[i+1:])
			if l > r {
				return l
			}
			return r
		}

		if nums[i] < 0 {
			negativeIndex = append(negativeIndex, i)
		}
	}

	if len(negativeIndex) == 0 || len(negativeIndex)%2 == 0 {
		return len(nums)
	}

	var max int
	rl := []int{negativeIndex[0], negativeIndex[len(negativeIndex)-1], len(nums) - 1 - negativeIndex[0], len(nums) - 1 - negativeIndex[len(negativeIndex)-1]}

	fmt.Println(negativeIndex, rl)

	for _, v := range rl {
		if v > max {
			max = v
		}
	}

	return max
}

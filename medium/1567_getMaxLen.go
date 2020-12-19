package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, -6, 4, 0, 10}
	fmt.Println(getMaxLen(nums))
}

func getMaxLen(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res int
	var negIdx []int
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
			negIdx = append(negIdx, i)
		}
	}

	if len(negIdx)%2 == 0 {
		return len(nums)
	}

	if len(negIdx) > 0 {
		resArr := []int{negIdx[0], negIdx[len(negIdx)-1],
			len(nums) - 1 - negIdx[0], len(nums) - 1 - negIdx[len(negIdx)-1]}

		fmt.Println(negIdx, resArr)

		for _, v := range resArr {
			if v > res {
				res = v
			}
		}
	}

	return res
}

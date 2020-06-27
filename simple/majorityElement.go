package main

func main() {
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	println(majorityElement(nums))
}

// majorityElement
// 数组中占比超过一半的元素称之为主要元素
// 给定一个整数数组，找到它的主要元素。若没有，返回-1
func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	mapTemp := make(map[int]int) // map[value][times]

	for _, v := range nums {
		if mapTempV, ok := mapTemp[v]; ok {
			if mapTempV >= length/2 {
				return v
			}
			mapTemp[v]++
		} else {
			mapTemp[v] = 1
		}
	}

	return -1
}

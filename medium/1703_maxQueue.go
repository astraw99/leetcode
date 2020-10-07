package main

import "fmt"

func main() {
	q := Constructor()
	fmt.Println(q.Max_value())
}

type MaxQueue struct {
	Data []int
	Max  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		Data: nil,
		Max:  nil,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.Data = append(this.Data, value)
	if len(this.Max) == 0 {
		this.Max = append(this.Max, value)
	} else {
		if value >= this.Max[0] {
			this.Max = append([]int{}, value)
			return
		}

		for value > this.Max[len(this.Max)-1] {
			this.Max = this.Max[:len(this.Max)-1]
		}
		this.Max = append(this.Max, value)
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Data) == 0 {
		return -1
	}

	v := this.Data[0]
	this.Data = this.Data[1:]

	if v == this.Max[0] {
		this.Max = this.Max[1:]
	}

	return v
}

func binaryAdd(nums []int, v int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right - left) / 2
		if v > nums[mid] {
			left = mid + 1 + left
		} else {
			right = mid - 1
		}
	}

	if v > nums[left] {
		nums = append(nums[:left], v)
		nums = append(nums, nums[left:]...)
	} else {
		nums = append(nums[:left+1], v)
		nums = append(nums, nums[left+1:]...)
	}

	return nums
}

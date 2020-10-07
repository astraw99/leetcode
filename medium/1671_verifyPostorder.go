package main

import "fmt"

func main() {
	postOrder := []int{5, 2, -17, -11, 25, 76, 62, 98, 92, 61}
	fmt.Println(verifyPostorder(postOrder))
	fmt.Println(verifyPostorder2(postOrder))
	fmt.Println(verifyPostorder_(postOrder))

	fmt.Println(int(^uint(0) >> 1))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	sep := len(postorder) - 1
	root := postorder[sep]
	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= root {
			sep = i
		}
		if i > sep && postorder[i] < root {
			return false
		}
	}

	return verifyPostorder(postorder[:sep]) && verifyPostorder(postorder[sep:len(postorder)-1])
}

func verifyPostorder2(postorder []int) bool {
	return helper(postorder, 0, len(postorder)-1)
}

func helper(postOrder []int, left, right int) bool {
	if left >= right {
		return true
	}

	mid := left
	root := postOrder[right]
	for postOrder[mid] < root {
		mid++
	}
	for i := mid; i < right; i++ {
		if postOrder[i] < root {
			return false
		}
	}

	return helper(postOrder, left, mid-1) && helper(postOrder, mid, right-1)
}

// self wrong!!
func verifyPostorder_(postorder []int) bool {
	var res bool
	if len(postorder) == 0 {
		return true
	}

	rootVal := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	var leftPos, leftData []int
	var rightPos, rightData []int

	for i, _ := range postorder {
		if postorder[i] < rootVal {
			if len(leftPos) > 0 {
				if i-leftPos[len(leftPos)-1] != 1 {
					return false
				}
			}
			leftPos = append(leftPos, i)
			leftData = append(leftData, postorder[i])
		} else {
			if len(rightPos) > 0 {
				if i-rightPos[len(rightPos)-1] != 1 {
					return false
				}
			}
			rightPos = append(rightPos, i)
			rightData = append(rightData, postorder[i])
		}
	}

	res = true

	if len(leftData) > 0 {
		if len(leftData) >= 3 {
			if leftData[0] > leftData[1] || leftData[0] > leftData[2] {
				return false
			}
		}
		res = verifyPostorder_(leftData)
	}
	if res && len(rightData) > 0 {
		if len(rightData) >= 3 {
			if rightData[0] > rightData[1] || rightData[0] > rightData[2] {
				return false
			}
		}
		res = verifyPostorder_(rightData)
	}

	return res
}

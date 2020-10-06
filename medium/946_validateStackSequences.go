package main

import "fmt"

func main() {
	pushed := []int{1, 0}
	popped := []int{1, 0} // 4,3,5,1,2
	//fmt.Println(validateStackSequences_(pushed, popped))
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}

	var tempStack []int
	tempStack = append(tempStack, pushed[0])
	pushed = pushed[1:]

	for {
		if len(tempStack) > 0 && tempStack[len(tempStack)-1] == popped[0] {
			tempStack = tempStack[:len(tempStack)-1]
			popped = popped[1:]
		} else if len(pushed) > 0 {
			tempStack = append(tempStack, pushed[0])
			pushed = pushed[1:]
		} else {
			return false
		}

		if len(popped) == 0 {
			return true
		}
	}
}

// self wrong!!
func validateStackSequences_(pushed []int, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}

	if pushed[len(pushed)-1] == popped[0] {
		pushed = pushed[:len(pushed)-1]
		popped = popped[1:]
		return validateStackSequences_(pushed, popped)
	}

	pushed, unpushed := subVal(pushed, popped[0])
	popped = popped[1:]
	return validateStackSequences_(pushed, popped) || validateStackSequences_(append(pushed, unpushed...), popped)

	/*var res bool
	if len(pushed) != len(popped) {
		return false
	}

	l := len(pushed)
	count := 0
	for i := l - 1; i >= 0; i-- {
		if pushed[i] == popped[l-1-i] {
			count++
		} else {
			pushed, unpushed := subVal(pushed, popped[0])
			popped = popped[1:]
			return validateStackSequences(pushed, popped) || validateStackSequences(append(pushed, unpushed...), popped)
		}
	}

	if count == l {
		return true
	}

	return res*/
}

func subVal(pushed []int, v int) (p, up []int) {
	var unpushed []int

	for i, _ := range pushed {
		if pushed[i] == v {
			pushed = pushed[:i]
			if i+1 <= len(pushed)-1 {
				unpushed = pushed[i+1:]
			}
			break
		}
	}

	return pushed, unpushed
}

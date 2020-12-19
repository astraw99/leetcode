package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue: nil}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	return
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.queue) > 0 {
		last := this.queue[len(this.queue)-1]
		this.queue = this.queue[:len(this.queue)-1]
		return last
	}

	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queue) > 0 {
		return this.queue[len(this.queue)-1]
	}

	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

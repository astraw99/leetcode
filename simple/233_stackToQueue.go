package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stack1: nil, stack2: nil}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack2 = append(this.stack2, x)
	return
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stack1) == 0 {
		for _, v := range this.stack2 {
			this.stack1 = append(this.stack1, v)
		}
		this.stack2 = nil
	}

	first := this.stack1[0]
	this.stack1 = this.stack1[1:]

	return first
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stack1) == 0 {
		for _, v := range this.stack2 {
			this.stack1 = append(this.stack1, v)
		}
		this.stack2 = nil
	}
	if len(this.stack1) > 0 {
		return this.stack1[0]
	}

	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

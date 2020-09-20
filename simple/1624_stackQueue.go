package main

import "fmt"

func main() {
	obj := Constructor()
	val0 := obj.DeleteHead()
	obj.AppendTail(5)
	obj.AppendTail(2)
	//obj.AppendTail(30)
	//val := obj.DeleteHead()
	//val2 := obj.DeleteHead()
	val3 := obj.DeleteHead()
	val4 := obj.DeleteHead()

	fmt.Println(val0, val3, val4, obj.NormalStack, obj.ReversedStack)
}

// 用两个栈实现一个队列。
// 队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead 操作返回 -1 )
type CQueue struct {
	NormalStack   []int
	ReversedStack []int
}

func Constructor() CQueue {
	return CQueue{
		NormalStack:   nil,
		ReversedStack: nil,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.NormalStack = append(this.NormalStack, value)
	length := len(this.NormalStack)

	fmt.Println(this.NormalStack)

	this.ReversedStack = nil
	for i := length - 1; i >= 0; i-- {
		this.ReversedStack = append(this.ReversedStack, this.NormalStack[i])
	}

	fmt.Println(this.ReversedStack)
}

func (this *CQueue) DeleteHead() int {
	length := len(this.ReversedStack)
	if length == 0 {
		return -1
	}

	this.NormalStack = this.NormalStack[1:] // 去掉第一个元素

	res := this.ReversedStack[length-1]
	this.ReversedStack = this.ReversedStack[:length-1]

	fmt.Println(this.NormalStack, this.ReversedStack)

	return res
}

package main

import (
	"fmt"
)

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)                        // 缓存是 {1=1}
	obj.Put(2, 2)                        // 缓存是 {1=1, 2=2}
	fmt.Println(obj.Get(1), obj.MapData) // 返回 1
	obj.Put(3, 3)                        // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(obj.Get(2))              // 返回 -1 (未找到)
	obj.Put(4, 4)                        // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(obj.Get(1))              // 返回 -1 (未找到)
	fmt.Println(obj.Get(3))              // 返回 3
	fmt.Println(obj.Get(4))              // 返回 4
}

type LRUCache struct {
	Cap        int
	Head, Tail *ListNode         // dummy node
	MapData    map[int]*ListNode // map[key][node]
}

type ListNode struct {
	Key, Val   int
	Prev, Next *ListNode
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{}
	tail := &ListNode{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Cap:     capacity,
		Head:    head,
		Tail:    tail,
		MapData: make(map[int]*ListNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.MapData[key]; ok {
		this.Remove(node)
		this.AddFront(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// exist
	if node, ok := this.MapData[key]; ok {
		node.Val = value
		this.Remove(node)
		this.AddFront(node)
		return
	}

	// not exist
	if len(this.MapData) >= this.Cap {
		// first to del map
		delete(this.MapData, this.Tail.Prev.Key)
		// then to remove node
		this.Remove(this.Tail.Prev)
	}

	newNode := &ListNode{
		Key: key,
		Val: value,
	}
	this.AddFront(newNode)
	this.MapData[key] = newNode
}

func (this *LRUCache) Remove(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddFront(node *ListNode) {
	// first to connect next
	node.Next = this.Head.Next
	this.Head.Next.Prev = node

	// then to connect head
	this.Head.Next = node
	node.Prev = this.Head
}

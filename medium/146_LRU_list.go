package main

import (
	"container/list"
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
	Cap     int
	List    *list.List            // []int{key}
	MapData map[int]*list.Element // map[key][val]
}

type Node struct {
	Key int
	Val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:     capacity,
		List:    list.New(),
		MapData: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.MapData[key]; ok {
		this.List.MoveToFront(node)
		return node.Value.(Node).Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// exist
	if node, ok := this.MapData[key]; ok {
		this.List.MoveToFront(node)
		node.Value = Node{key, value}
		return
	}

	// not exist
	if this.List.Len() >= this.Cap {
		d := this.List.Back()
		this.List.Remove(d)
		delete(this.MapData, d.Value.(Node).Key)
	}

	this.MapData[key] = this.List.PushFront(Node{key, value})
}

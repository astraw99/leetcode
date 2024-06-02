package main

import "fmt"

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
	Keys    []int       // []int{key}
	MapData map[int]int // map[key][val]
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:     capacity,
		Keys:    nil,
		MapData: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.MapData[key]; ok {
		this.Keys = DelKey(this.Keys, key)
		this.Keys = append([]int{key}, this.Keys...)
		return val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// exist
	if _, ok := this.MapData[key]; ok {
		this.Keys = DelKey(this.Keys, key)
		this.Keys = append([]int{key}, this.Keys...)

		this.MapData[key] = value
		return
	}

	// not exist
	if len(this.MapData) >= this.Cap {
		delKey := this.Keys[len(this.Keys)-1]
		this.Keys = DelKey(this.Keys, delKey)
		delete(this.MapData, delKey)
	}

	this.Keys = append([]int{key}, this.Keys...)
	this.MapData[key] = value
}

func DelKey(keys []int, key int) []int {
	for i := 0; i < len(keys); i++ {
		if keys[i] == key {
			return append(keys[:i], keys[i+1:]...)
		}
	}

	return keys[:len(keys)-1]
}

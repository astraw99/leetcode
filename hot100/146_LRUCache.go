/*
 * Copyright (C) 2023 https://github.com/astraw99. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License"); you may not use this
 * file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

func main() {

}

type LRUCache struct {
	keys     []int
	data     map[int]int // key -> val
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keys:     make([]int, 0),
		data:     make(map[int]int),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.data[key]; ok {

		for i, k := range this.keys {
			if k == key {
				this.keys = append(this.keys[:i], this.keys[i+1:]...)
				this.keys = append([]int{k}, this.keys...)
				break
			}
		}

		return val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		for i, k := range this.keys {
			if k == key {
				this.keys = append(this.keys[:i], this.keys[i+1:]...)
				this.keys = append([]int{k}, this.keys...)
				break
			}
		}
		this.data[key] = value
	} else {
		if len(this.keys) == this.capacity {
			delete(this.data, this.keys[len(this.keys)-1])
			this.keys = this.keys[:len(this.keys)-1]
		}

		this.keys = append([]int{key}, this.keys...)
		this.data[key] = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

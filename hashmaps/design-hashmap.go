package hashmap

import (
	"fmt"
)

const (
	mod = 1024
)

type MyHashMap struct {
	set [mod]*listNode
}

type listNode struct {
	key  int
	val  int
	next *listNode
}

func Constructor() MyHashMap {
	return MyHashMap{set: [mod]*listNode{}}
}

func (this *MyHashMap) Get(key int) int {
	ptr := this.set[key%mod]
	for ptr != nil {
		if ptr.key == key {
			return ptr.val
		} else {
			ptr = ptr.next
		}
	}
	return -1
}

func (this *MyHashMap) Put(key int, val int) {
	i := key % mod
	ptr := this.set[i]
	for ptr != nil {
		if ptr.key == key {
			ptr.val = val
			return
		} else {
			ptr = ptr.next
		}
	}
	node := &listNode{key: key, val: val, next: this.set[i]}
	this.set[i] = node
}

func (this *MyHashMap) Remove(key int) {
	i := key % mod
	ptr := this.set[i]
	prev := &listNode{next: ptr}
	head := prev
	for ptr != nil {
		if ptr.key == key {
			prev.next = ptr.next
			break
		} else {
			prev = prev.next
			ptr = ptr.next
		}
	}
	this.set[i] = head.next
}

func DesignHashMapTest() {
	// 	["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
	// 	[[], 		 [1, 1], [2, 2], [1],   [3],   [2, 3], [2],  [2], 	   [2]]
	fmt.Println("\nDESIGN HASH MAP using linked list and array:")
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 3)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))
}

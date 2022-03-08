package hashmap

import (
	"fmt"
)

func DesignHashMap() {
	fmt.Println("\nExercise for designing a hash map using linked list and array:")
}

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

package lrucache

import "fmt"

type LRUCache struct {
	capacity int
	store    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, store: make(map[int]*Node)}
}

func (lru *LRUCache) AddToStore(node *Node) {
	node.prev = nil
	node.next = lru.head
	if node.next != nil {
		node.next.prev = node
	}
	lru.head = node
	if lru.tail == nil {
		lru.tail = node
	}
}

func (lru *LRUCache) RemoveFromStore(node *Node) {
	if node == lru.head {
		lru.head = node.next
	}
	if node == lru.tail {
		lru.tail = node.prev
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
}

func (lru *LRUCache) Put(key int, val int) {
	node, ok := lru.store[key]
	if ok {
		lru.RemoveFromStore(node)
		node.val = val
	} else {
		node = &Node{key: key, val: val}
		lru.store[key] = node
	}

	lru.AddToStore(node)
	if len(lru.store) > lru.capacity {
		node = lru.tail
		if node != nil {
			lru.RemoveFromStore(node)
			delete(lru.store, node.key)
		}
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.store[key]
	if ok {
		lru.RemoveFromStore(node)
		lru.AddToStore(node)
		return node.val
	}
	return -1
}

func TestLRUCache() {

	fmt.Println("\nLRU Cache:")
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println("Expected: 1", obj.Get(1))
	obj.Put(3, 3)
	fmt.Println("Expected: -1", obj.Get(2))
	obj.Put(4, 4)
	fmt.Println("Expected: -1", obj.Get(1))
	fmt.Println("Expected: 3", obj.Get(3))
	fmt.Println("Expected: 4", obj.Get(4))

}

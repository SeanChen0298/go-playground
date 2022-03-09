package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type SinglyLinkedList struct {
	len  int
	head *ListNode
}

func initList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (s *SinglyLinkedList) ConstructLinkedList(nums []int) {
	for _, val := range nums {
		node := &ListNode{}
		node.Val = val
		if s.head == nil {
			s.head = node
		} else {
			node.Next = s.head
			s.head = node
		}
		s.len++
	}
}

func (s *SinglyLinkedList) Display() {
	var result []int
	current := s.head
	for current.Next != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	result = append(result, current.Val)
	fmt.Println(result)
}

func TestConstructLinkedList() {
	nums := []int{1, 2, 3}
	l1 := initList()
	l1.ConstructLinkedList(nums)
	l1.Display()
}

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
	current := s.head
	for _, val := range nums {
		node := &ListNode{Val: val, Next: nil}
		if s.head == nil {
			s.head = node
			current = s.head
		} else {
			current.Next = node
			current = current.Next
		}
		s.len++
	}
}

func (s *SinglyLinkedList) Display() {
	current := s.head
	for current.Next != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Print(current.Val)
}

func Display(head *ListNode) {
	current := head
	for current.Next != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Print(current.Val)
}

func TestConstructLinkedList() {
	fmt.Println("\nArray to LinkedList:")
	nums := []int{1, 2, 3}
	l1 := initList()
	l1.ConstructLinkedList(nums)
	l1.Display()
}

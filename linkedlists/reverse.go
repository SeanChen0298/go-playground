package linkedlist

import "fmt"

func (s *SinglyLinkedList) Reverse() {
	_, root := reverseRecursive(s.head)
	s.head = root
}
func reverseRecursive(head *ListNode) (cur, root *ListNode) {
	if head == nil {
		return nil, nil
	}

	prev, root := reverseRecursive(head.Next)
	if prev == nil {
		return head, head
	}

	// Actual Reverse
	cur = &ListNode{Val: head.Val}
	prev.Next = cur
	return cur, root
}

func TestReverse() {

	nums := []int{1, 2, 3, 4, 5}
	l1 := initList()
	l1.ConstructLinkedList(nums)
	fmt.Println("\nOriginal LinkedList:")
	l1.Display()

	fmt.Println("\nReversed LinkedList:")
	l1.Reverse()
	l1.Display()
}

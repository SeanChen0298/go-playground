package linkedlist

import "fmt"

func MergeSortedLinkedList(l1, l2 *ListNode) *ListNode {
	merged := ListNode{}
	merged.Next = nil
	current := &merged

	for l1 != nil || l2 != nil {
		if l1 == nil {
			(*current).Next = l2
			break
		} else if l2 == nil {
			(*current).Next = l1
			break
		} else {
			if l1.Val <= l2.Val {
				(*current).Next = l1
				l1 = (*l1).Next
			} else {
				(*current).Next = l2
				l2 = (*l2).Next
			}
		}
		current = current.Next
	}
	return merged.Next
}

func TestMergeLinkedList() {
	fmt.Println("\nMerge Sorted LinkedList:")
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4, 5, 6}

	fmt.Println("\n1st LinkedList:")
	l1 := initList()
	l1.ConstructLinkedList(nums1)
	l1.Display()

	fmt.Println("\n2nd LinkedList:")
	l2 := initList()
	l2.ConstructLinkedList(nums2)
	l2.Display()

	fmt.Println("\nMerged LinkedList:")
	merged := MergeSortedLinkedList(l1.head, l2.head)
	Display(merged)
}

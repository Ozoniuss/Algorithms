package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// The idea here is to use two pointers. Both pointers start at the first node.
// We'll define the "left" and "right" pointers as follows:
// - If both point to the same node, either of them can be the
//   left or right pointer;
// - If they do not point to the same node, we call the left pointer the pointer
//   p that can become the other pointer after a succession of operations
//   p = p.Next;
// - If at any point in time there is a single non-nil pointer pointing to a node,
//   that is going to be the left pointer.
//
// Start with both pointers in the same location, and at each step we move
// right to the current node's Next node.
// If the value pointed by right is the same as the value pointed by left,
// we can repeat this operation until this condition is no longer true.
// Once that happens, all the nodes from left to right (non-including right)
// have the same value. We can set left's node's .Next value to point to
// right's node, and we can set left to right.
// When right points to nil, we can set left's node's .Next value to nil.

// deleteDuplicates removes the duplicate elements from a sorted linked list.
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	previous := head
	current := head
	for current != nil {
		if current.Val == previous.Val {
			current = current.Next
		} else {
			previous.Next = current
			previous = current
		}
	}
	previous.Next = nil
	return head
}

func printLinked(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}

func convertLinkedToArray(head *ListNode) []int {
	elems := []int{}
	for head.Next != nil {
		elems = append(elems, head.Val)
	}
	return elems
}

func convertLinkedFromArray(elems []int) *ListNode {
	if len(elems) == 0 {
		return nil
	}
	if len(elems) == 1 {
		return &ListNode{Val: elems[0], Next: nil}
	}
	head := &ListNode{Val: elems[0]}
	prev := head
	for i := 1; i < len(elems); i++ {
		node := &ListNode{Val: elems[i], Next: nil}
		prev.Next = node
		prev = prev.Next
	}
	return head
}

func main() {
	duplicates := convertLinkedFromArray([]int{4, 4, 3, 3, 2, 2, 1, 1})
	b := deleteDuplicates(duplicates)
	printLinked(b)
}

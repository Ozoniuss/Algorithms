package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	sz := 0

	cur := head
	for cur != nil {
		cur = cur.Next
		sz += 1
	}

	// we need to remove this element from the start of the list
	n = sz - n + 1

	// edge case of removing first element
	if n == 1 {
		head = head.Next
		return head
	}

	// basically need to remove cur
	i := 1
	cur = head
	prev := new(ListNode)
	for i < n {
		prev = cur
		cur = cur.Next
		i++
	}
	if cur == nil {
		panic("not enough elements")
	}
	if cur.Next == nil {
		prev.Next = nil
	} else {
		prev.Next = cur.Next
	}

	return head
}

func main() {

}

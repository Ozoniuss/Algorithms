package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	b := &strings.Builder{}
	cur := l
	for cur != nil {
		b.WriteString(fmt.Sprintf("%d,", cur.Val))
		cur = cur.Next
	}
	s := b.String()
	return s[:len(s)-1]

}

// algorithm for reversing regular linked list, returning a pointer to
// the head
func reverseRegular(head *ListNode) (*ListNode, *ListNode) {
	cur := head
	prev := (*ListNode)(nil)
	for cur != nil {
		fmt.Println(cur.Val)
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	cur := head
	sz := 1

	for sz < k && cur != nil {
		sz += 1
		cur = cur.Next
	}

	// we just need to reverse these elements because the list doesn't
	// actually have enough elements
	if cur == nil {
		// h, _ := reverseRegular(head)
		// return h
		return head
	}

	// we need to reverse these elements only, and then extend the reversed
	// list's tail with the next reversed list
	//
	// note that cur is the last element
	nextlist := cur.Next
	cur.Next = nil
	h, t := reverseRegular(head)
	t.Next = reverseKGroup(nextlist, k)
	return h
}

func main() {
	l6 := &ListNode{
		Val:  6,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: l6,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	fmt.Println(l1)

	fmt.Println(reverseKGroup(l1, 4))

	// r1, r2 := reverseRegular(l1)
	// fmt.Println(r1, r2)

}

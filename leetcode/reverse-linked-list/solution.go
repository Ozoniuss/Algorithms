package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		last := head.Next
		last.Next = head
		head.Next = nil
		return last
	}

	previous := head
	current := head.Next
	head.Next = nil

	for current != nil {
		future := current.Next
		current.Next = previous
		previous = current
		current = future
	}

	return previous
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
	numbers := convertLinkedFromArray([]int{1, 2, 3, 4, 5})
	h := reverseList(numbers)
	printLinked(h)
}

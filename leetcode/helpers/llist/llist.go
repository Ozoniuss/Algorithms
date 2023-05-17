package llist

import (
	"fmt"
	"strconv"
	"strings"
)

// Leetcode definition for singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type ErrInvalidInputArray struct {
	msg string
}

func NewErrInvalidInputArray(msg string) ErrInvalidInputArray {
	return ErrInvalidInputArray{
		msg: msg,
	}
}

func NewErrInvalidInputArrayf(format string, value ...any) ErrInvalidInputArray {
	return ErrInvalidInputArray{
		msg: fmt.Sprintf(format, value...),
	}
}

func (e ErrInvalidInputArray) Error() string {
	return e.msg
}

// FromStringArray converts a leetcode specific string array to a linked list.
// Sample input: [5,4,2,1]
func FromStringArray(in string) (*ListNode, error) {
	if in == "" || in == "[]" {
		return nil, nil
	}

	if len(in) < 2 {
		return nil, NewErrInvalidInputArrayf("invalid input: %s", in)
	}

	if in[0] != '[' && in[len(in)-1] != ']' {
		return nil, NewErrInvalidInputArrayf("invalid input: %s", in)
	}

	// Trim parentheses
	in = in[1 : len(in)-1]

	curr := 0
	prev := 0
	head := &ListNode{
		Next: nil,
	}
	l := &ListNode{
		Next: head,
	}
	for curr <= len(in) {
		// Lazy evaluation helps to not panic.
		if curr == len(in) || in[curr] == ',' {
			l = l.Next
			number := in[prev:curr]
			val, err := strconv.Atoi(number)
			if err != nil {
				return nil, NewErrInvalidInputArrayf("invalid input element: %s", number)
			}
			l.Val = val
			l.Next = &ListNode{
				Next: nil,
			}
			prev = curr + 1
		}
		curr++
	}
	l.Next = nil
	return head, nil
}

// Equal checks that two ListNodes have the same elements.
func Equal(l1, l2 *ListNode) bool {

	if (l1 != l2) && ((l1 == nil) || (l2 == nil)) {
		return false
	}

	cur1 := l1
	cur2 := l2

	for cur1 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur2 == nil
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	out := strings.Builder{}
	cur := l
	for cur != nil {
		out.WriteString(strconv.Itoa(cur.Val))
		if cur.Next != nil {
			out.WriteString(" -> ")
		}
		cur = cur.Next
	}
	return out.String()
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	ret := []int{-1, -1}

	h := head.Next
	p := head
	fpos := -1
	lpos := -1
	i := 0

	for h.Next != nil {
		i += 1
		hh := h
		pp := p
		h = h.Next
		p = p.Next
		if !((hh.Val > pp.Val && hh.Val > hh.Next.Val) || (hh.Val < pp.Val && hh.Val < hh.Next.Val)) {
			continue
		}

		// first critical point
		if fpos == -1 {
			fpos = i
			lpos = i
			continue
		}

		// at this point there are at least two points
		ret[0] = min(ret[0], i-lpos)
		if ret[0] == -1 {
			ret[0] = i - lpos
		}
		ret[1] = i - fpos
		lpos = i
	}

	return ret
}

func main() {
	h := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(nodesBetweenCriticalPoints(h))
}

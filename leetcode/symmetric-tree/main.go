package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isReflectiveTree(p *TreeNode, q *TreeNode) bool {

	// Trivial cases
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	if !isReflectiveTree(p.Left, q.Right) {
		return false
	}
	if !isReflectiveTree(p.Right, q.Left) {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isReflectiveTree(root.Left, root.Right)
}

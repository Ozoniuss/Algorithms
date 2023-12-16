package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	_, ok := findMaximumIfValid(root, -1<<31-1)
	return ok
}

// findMaximumIfValid returns the maximum value of a BST, if the BST is valid.
// The idea is that a depth-first walk should walk through the numbers in
// ascending order.
func findMaximumIfValid(root *TreeNode, prev int) (int, bool) {
	if root == nil {
		return prev, true
	}
	maxNumber, ok := findMaximumIfValid(root.Left, prev)
	if !ok {
		return 0, false
	}
	if maxNumber >= root.Val {
		return 0, false
	}
	return findMaximumIfValid(root.Right, root.Val)
}

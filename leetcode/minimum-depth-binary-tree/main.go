package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + dfs(root)
}

func dfs(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return 0
	}
	if node.Left == nil && node.Right != nil {
		return 1 + dfs(node.Right)
	}
	if node.Left != nil && node.Right == nil {
		return 1 + dfs(node.Left)
	}
	return 1 + min(dfs(node.Left), dfs(node.Right))
}

func main() {

}

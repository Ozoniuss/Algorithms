package main

import (
	"fmt"
)

func inorderTraversal(root *TreeNode) []int {
	visited := make([]int, 0)
	traverse(root, &visited)
	return visited
}

func traverse(node *TreeNode, visited *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, visited)
	*visited = append(*visited, node.Val)
	traverse(node.Right, visited)
}

func main() {
	root, err := ParseTree("tree.txt")
	if err != nil {
		fmt.Println(err)
	}
	vals := inorderTraversal(root)
	fmt.Println(vals)

	vals = inorderTraversal(nil)
	fmt.Println(vals)
}

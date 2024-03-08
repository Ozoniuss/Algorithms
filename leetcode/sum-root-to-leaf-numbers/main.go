package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sum("", root)
}

func sum(curr string, root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		v, _ := strconv.Atoi(curr + strconv.Itoa(root.Val))
		fmt.Println(v)
		return v
	}
	if root.Left == nil {
		return sum(curr+strconv.Itoa(root.Val), root.Right)
	}
	if root.Right == nil {
		return sum(curr+strconv.Itoa(root.Val), root.Left)
	}
	return sum(curr+strconv.Itoa(root.Val), root.Left) + sum(curr+strconv.Itoa(root.Val), root.Right)
}

func main() {

}

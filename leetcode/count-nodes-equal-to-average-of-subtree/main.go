package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	total := 0
	helper(root, &total)
	return total
}

func helper(root *TreeNode, total *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	// leaf always counts
	if root.Left == nil && root.Right == nil {
		*total += 1
		return 1, root.Val
	}

	cl, tl := helper(root.Left, total)
	cr, tr := helper(root.Right, total)

	avg := (tl + tr + root.Val) / (cl + cr + 1)
	if avg == root.Val {
		*total += 1
	}
	return cl + cr + 1, tl + tr + root.Val
}

func main() {

}

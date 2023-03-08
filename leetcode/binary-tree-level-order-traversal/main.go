package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	ret := make([][]int, 0)
	for len(q) > 0 {

		level := make([]int, 0)
		tempq := make([]*TreeNode, 0)

		// Read entire level
		for len(q) > 0 {
			top := q[0]
			q = q[1:]

			level = append(level, top.Val)

			if top.Left != nil {
				tempq = append(tempq, top.Left)
			}
			if top.Right != nil {
				tempq = append(tempq, top.Right)
			}
		}
		q = tempq
		ret = append(ret, level)
	}
	return ret
}

func main() {

}

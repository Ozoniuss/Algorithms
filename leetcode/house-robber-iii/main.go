package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type cachekey struct {
	node *TreeNode
	skip bool
}

func rob(root *TreeNode) int {
	return max(robh(root, true, make(map[cachekey]int)), robh(root, false, make(map[cachekey]int)))
}

func robh(node *TreeNode, skip bool, cache map[cachekey]int) int {
	if node == nil {
		return 0
	}

	if v, ok := cache[cachekey{
		node: node,
		skip: skip,
	}]; ok {
		return v
	}

	// case when it is nil and skip will always return 0
	if !skip && node.Left == nil && node.Right == nil {
		cache[cachekey{
			node: node,
			skip: skip,
		}] = node.Val
		return node.Val
	}

	if !skip {
		// I must skip both left and right in this case
		v := node.Val + robh(node.Left, true, cache) + robh(node.Right, true, cache)
		cache[cachekey{
			node: node,
			skip: skip,
		}] = v
		return v
	} else {
		// I can take left or right or skip either
		v1 := robh(node.Left, false, cache) + robh(node.Right, true, cache)
		v2 := robh(node.Left, true, cache) + robh(node.Right, false, cache)
		v3 := robh(node.Left, false, cache) + robh(node.Right, false, cache)
		v4 := robh(node.Left, true, cache) + robh(node.Right, true, cache)
		cache[cachekey{
			node: node,
			skip: skip,
		}] = max(v1, v2, v3, v4)
		return max(v1, v2, v3, v4)
	}
}

func main() {

}

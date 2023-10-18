class TreeNode {
	val: number;
	left: TreeNode | null;
	right: TreeNode | null;
	constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
		this.val = val === undefined ? 0 : val;
		this.left = left === undefined ? null : left;
		this.right = right === undefined ? null : right;
	}
}

function traverse(current: TreeNode | null, explored: number[]) {
	if (!current) {
		return;
	}
	explored.push(current.val);
	traverse(current.left, explored);
	traverse(current.right, explored);
}

function preorderTraversal(root: TreeNode | null): number[] {
	if (!root) {
		return [];
	}
	const explored = [];
	traverse(root, explored);
	return explored;
}

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

function height(treeNode: TreeNode | null): number {
	if (!treeNode) {
		return 0;
	}
	if (!treeNode.left && !treeNode.right) {
		return 1;
	}
	if (!treeNode.left) {
		return 1 + height(treeNode.right);
	}
	if (!treeNode.right) {
		return 1 + height(treeNode.left);
	}
	return (
		1 +
		(height(treeNode.left) > height(treeNode.right)
			? height(treeNode.left)
			: height(treeNode.right))
	);
}

function isBalanced(root: TreeNode | null): boolean {
	if (!root) {
		return true;
	}
	return (
		isBalanced(root.left) &&
		isBalanced(root.right) &&
		Math.abs(height(root.left) - height(root.right)) <= 1
	);
}

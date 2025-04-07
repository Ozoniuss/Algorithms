# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:

    paths = []

    def pathSum(self, root, targetSum):
        """
        :type root: Optional[TreeNode]
        :type targetSum: int
        :rtype: List[List[int]]
        """
        path = []
        self.paths = []
        self.explore(root, path, 0, targetSum)
        return self.paths

    def explore(self, node: TreeNode, path: list, current, target_sum):
        if node == None:
            return
        path.append(node.val)
        if (
            node.left == None
            and node.right == None
            and node.val + current == target_sum
        ):

            self.paths.append(path.copy())

        if node.left != None:
            self.explore(node.left, path, current + node.val, target_sum)

        if node.right != None:
            self.explore(node.right, path, current + node.val, target_sum)

        path.pop()

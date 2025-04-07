# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def hasPathSum(self, root, targetSum):
        """
        :type root: Optional[TreeNode]
        :type targetSum: int
        :rtype: bool
        """
        return self.check_sum(root, 0, targetSum)

    def check_sum(self, node: TreeNode, current, target):
        if node == None:
            return False
        if node.left == None and node.right == None:
            return node.val + current == target
        else:
            return self.check_sum(
                node.left, current + node.val, target
            ) or self.check_sum(node.right, current + node.val, target)

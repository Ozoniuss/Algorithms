from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if root == None:
            return True
        return self.findMaxNumberInTree(root, -(2**31) - 1) is not None

    def findMaxNumberInTree(self, root: Optional[TreeNode], prev: Optional[int]):
        if root is None:
            return prev

        maxNumber = self.findMaxNumberInTree(root.left, prev)
        if maxNumber is None:
            return None
        if maxNumber > root.val:
            return None
        return self.findMaxNumberInTree(root.right, root.val)

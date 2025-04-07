from typing import Optional, List
from collections import defaultdict, deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    number_of_paths = 0

    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        self.number_of_paths = 0
        current_sums = []
        self.explore(root, current_sums, targetSum)
        return self.number_of_paths

    def explore(
        self,
        node: Optional[TreeNode],
        current_sums: list[int],
        targetSum: int,
    ):
        if node is None:
            return

        nxt = 0
        if len(current_sums) == 0:
            nxt = node.val
        else:
            last = current_sums[-1]
            nxt = last + node.val
        current_sums.append(nxt)

        print(current_sums)

        if node.left is not None:
            self.explore(node.left, current_sums, targetSum)

        if node.right is not None:
            self.explore(node.right, current_sums, targetSum)

        if nxt == targetSum:
            self.number_of_paths += 1
        for idx, el in enumerate(current_sums):
            if nxt - el == targetSum and idx != len(current_sums) - 1:
                self.number_of_paths += 1

        current_sums.pop()


class SolutionImproved:
    number_of_paths = 0

    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        self.number_of_paths = 0
        current_sums = defaultdict(int)
        self.explore(root, current_sums, 0, targetSum)
        return self.number_of_paths

    def explore(
        self,
        node: Optional[TreeNode],
        current_sums: list[int],
        last_sum: int,
        targetSum: int,
    ):
        if node is None:
            return

        nxt = last_sum + node.val

        if nxt == targetSum:
            self.number_of_paths += 1
        self.number_of_paths += current_sums[nxt - targetSum]

        current_sums[nxt] += 1

        if node.left is not None:
            self.explore(node.left, current_sums, nxt, targetSum)

        if node.right is not None:
            self.explore(node.right, current_sums, nxt, targetSum)

        current_sums[nxt] -= 1


# d = deque()

# d.append(7)
# d.append(8)
# print(d[-1])
# for el in enumerate(d):
#     print(el)

"""
# Definition for a Node.
"""
from typing import List

class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children

class Solution:
    def levelOrder(self, root: Node) -> List[List[int]]:
        if root is None:
            return []
        
        traversal = []
        q = [[root]]
        while len(q) != 0:
            top = q.pop(0)
            traversal.append([node.val for node in top])
            nextNodes = []
            for node in top:
                if node.children is None:
                    continue
                for next in node.children:
                    nextNodes.append(next)
            if len(nextNodes) != 0:
                q.append(nextNodes)
        
        return traversal

n = Node(val=1, children=[Node(val=2), Node(val=3), Node(val=4, children=[Node(val=5), Node(val=6)])])
s = Solution()
print(s.levelOrder(n))
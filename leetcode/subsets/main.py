from typing import List
from itertools import product

"""
Basically, BFS approach, where we create subsets (a1, a2, ..., aj) out of
(A_1, A_2, ..., A_n) such that a_i >= A_i. Each subset is a subsequence of 
the original sequence.

The queue contains the subsets that were already generated. So when we take the
first element of the queue, we find the first number that hadn't been added yet
(since a subset is basically just a subsequence), and we add that subset to the
queue.
"""
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        q = [[]]
        ret = []
        while len(q) != 0:
            top = q.pop(0)
            ret.append(top)
            if len(top) == len(nums):
                continue
            last = 0
            if len(top) != 0:
                last = nums.index(top[-1])+1
            for i in range(last, len(nums)):
                next = top.copy()
                next.append(nums[i])
                q.append(next)

        return ret
    
s = Solution()

print(s.subsets([1,2,3]))
print(s.subsets([0]))
print(s.subsets([3,2,4,1]))
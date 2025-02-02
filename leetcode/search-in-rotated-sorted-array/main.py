import re
from typing import List


# 4 5 6 7 8 9 10 0 1 2
# 100 101 3 4 5 6 7 8 9


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        first = nums[l]
        last = nums[r]

        if first == target:
            return l
        if last == target:
            return r

        while l < r:
            m = (l + r) // 2
            # print(l, m, r)
            num = nums[m]

            if num == target:
                return m
            
            elif num >= first: # [first, num) are ordered
                if first < target < num:
                    r = m-1
                else:
                    l = m+1
            
            elif num < first: # (first .... , smallest ... num ... last)
                if num < target < last:
                    l = m+1
                else:
                    r = m-1
        
        if l == r and nums[l] == target:
            return l

        return -1
    
s = Solution()
print(s.search([4,5,6,7,0,1,2], 0))
print(s.search([4,5,6,7,0,1,2], 3))
print(s.search([1], 0))
print(s.search([1, 3], 0))

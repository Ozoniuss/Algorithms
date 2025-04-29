# fully coded in leetcode editor
class Solution:
    def maximumGap(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return 0
        nums.sort()
        diff = 0
        for idx in range(len(nums) - 1):
            a = nums[idx]
            b = nums[idx+1]
            if b-a > diff:
                diff = b-a
        return diff


from typing import List

class Solution:
    def search(self, nums: List[int], target: int) -> int:

        if len(nums) == 0: return -1

        left_index = 0
        right_index = len(nums)

        while True:

            if left_index > right_index:
                return -1

            midpoint = (left_index + right_index) // 2

            if midpoint == len(nums):
                return -1

            if nums[midpoint] > target:
                right_index = midpoint - 1

            if nums[midpoint] < target:
                left_index = midpoint + 1

            if nums[midpoint] == target:
                return midpoint


a = Solution()
print(a.search(nums = [-1,0,3,5,9,12], target = 2))
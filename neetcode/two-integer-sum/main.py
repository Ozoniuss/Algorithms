class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        processed = {}
        for i, num in enumerate(nums):
            if target - num in processed:
                return [processed[target - num], i]
            processed[num] = i

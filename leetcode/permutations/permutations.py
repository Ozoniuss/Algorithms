from typing import List

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:

        permutations = [[]]

        for number in nums:
            new_permutations = []
            for p in permutations:
                preset_elements = set(p)
                for number2 in nums:
                    if number2 not in preset_elements:
                        new_permutations.append(p+[number2])
            permutations = new_permutations

        return permutations

s = Solution()
print(s.permute(nums = [1,2,3]))
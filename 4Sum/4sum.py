from typing import List

class Solution:
    def fourSumBruteForce(self, nums: List[int], target: int) -> List[List[int]]:

        # brute force is not optimal

        possibilities = []

        for a in range(0, len(nums)-3):
            for b in range(a+1, len(nums)-2):
                for c in range(b+1, len(nums)-1):
                    for d in range(c+1, len(nums)):
                        print(a,b,c,d)
                        if nums[a] + nums[b] + nums[c] + nums[d] == target:
                            sorted_nums = sorted([nums[a], nums[b], nums[c], nums[d]])
                            print(sorted_nums)
                            if len(possibilities) == 0:
                                possibilities.append([sorted_nums[0], sorted_nums[1], sorted_nums[2], sorted_nums[3]])
                            else:
                                present = False
                                for p in possibilities:
                                    if (p[0] == sorted_nums[0]) and (p[1] == sorted_nums[1]) and (p[2] == sorted_nums[2]) and (p[3] == sorted_nums[3]):
                                        present = True
                                if not present:
                                    possibilities.append([sorted_nums[0], sorted_nums[1], sorted_nums[2], sorted_nums[3]])


        return possibilities

s = Solution()

print(s.fourSum(nums = [-3,-2,-1,0,0,1,2,3], target = 0))
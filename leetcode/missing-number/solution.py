from typing import List

# Se solution.go for documentation

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        l = len(nums)
        sum = 0
        for _, num in enumerate(nums):
            sum += num
        
        return (l*(l+1)) // 2 - sum

s = Solution()
print(s.missingNumber([3,0,1]))
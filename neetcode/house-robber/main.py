import re
from typing import List
from xmlrpc.client import TRANSPORT_ERROR


class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return nums[0]
        dp = [(nums[0], True)]
        if nums[1] > nums[0]:
            dp.append((nums[1], True))
        else:
            dp.append((nums[0], False))

        for i in range(2, len(nums)):
            last, in_max = dp[i - 1]
            if not in_max:
                dp.append((nums[i] + last, True))
            else:
                last_before, _ = dp[i - 2]
                total = last_before + nums[i]
                if last > total:
                    dp.append((last, False))
                else:
                    dp.append((total, True))

        return dp[len(nums) - 1][0]


s = Solution()
print(s.rob([2, 9, 8, 3, 6]))

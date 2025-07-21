from functools import cache


class Solution:

    cache = {}

    def numTrees(self, n: int) -> int:

        if n == 0:
            return 1
        if n == 1:
            return 1
        
        if n in self.cache:
            return self.cache[n]

        s = 0
        for i in range(1, n+1):
            s += (self.numTrees(i-1) * self.numTrees(n-i))

        self.cache[n] = s    
    
        return s

s = Solution()
print(s.numTrees(3))
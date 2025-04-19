class Solution:
    cache = set()
    def __init__(self):
        self.cache = set()
    def numDecodings(self, s: str) -> int:
        if len(s) == 0:
            return 1
        if len(s) == 1:
            if s == '0':
                return 0
            return 1

        if s in self.cache:
            return self.cache[s]

        ab = s[:2]
        iab = int(s[:2])
        if ab[0] == '0':
            return 0
        
        if ab[1] == '0' and int(ab[0]) > 2:
            return 0
        
        if iab < 27:
            t = self.numDecodings(s[1:]) + self.numDecodings(s[2:])
            self.cache[s] = t
            return t
        t = self.numDecodings(s[1:])
        self.cache[s] = t
        return t
    
s = Solution()
print(s.numDecodings('226'))
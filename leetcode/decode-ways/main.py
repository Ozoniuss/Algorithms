class Solution:
    cache = {}
    def __init__(self):
        self.cache = {}
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

    def numDecodingsdp(self, s: str) -> int:
        dp = [0 for _ in range(len(s) + 1)]
        dp[-1] = 1

        if s[-1] == "0":
            dp[0] = 0
        else:
            dp[0] = 1 

        for i in range(1, len(s)):
            word = s[len(s)-1-i:]

            ab = word[:2]
            iab = int(word[:2])
            if ab[0] == '0':
                dp[i] = 0
        
            elif ab[1] == '0' and int(ab[0]) > 2:
                dp[i] = 0
        

            elif iab < 27:
                dp[i] = dp[i-1] + dp[i-2]
            else:
                dp[i] = dp[i-1]

        return dp[len(s)-1]


    
s = Solution()
print(s.numDecodings('226'))
print(s.numDecodingsdp('226'))
print(s.numDecodingsdp('10'))
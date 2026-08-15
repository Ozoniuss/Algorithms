from collections import defaultdict

class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        occurences = defaultdict(int)
        l = 0
        r = 0
        maxlen = -1
        while r < len(s):
            occurences[s[r]] += 1
            if occurences[s[r]] == 3:
                while occurences[s[r]] > 2:
                    occurences[s[l]] -= 1
                    l += 1
            maxlen = max(maxlen, r - l + 1)
            r += 1
        return maxlen

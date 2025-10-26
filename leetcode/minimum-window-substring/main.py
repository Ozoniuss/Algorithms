from collections import Counter

class Solution:
    def minWindow(self, s: str, t: str) -> str:

        if len(t) == 0:
            return ""

        if len(t) == 1:
            if t in s:
                return t
            else:
                return ""

        cnt = dict(Counter(t))

        lastocc = {}
        for k in cnt:
            lastocc[k] = []

        minwindow = 9999999999999999
        minstr = ""
        for (i, c) in enumerate(s):
            mpos = self.getShortestSol(lastocc, cnt,  c)
            if mpos == None:
                if c in lastocc:
                    lastocc[c].append(i)
                continue
            if i - mpos < minwindow:
                minwindow = i-mpos
                minstr = s[mpos:i+1]

            if c in lastocc:
                lastocc[c].append(i)
        
        return minstr
        # for c

    def getShortestSol(self,lastocc: dict, cnt: dict, cchar):
        (lastocc, cnt, cchar)
        if cchar not in lastocc:
            return None
        
        mpos = 9999999999999
        for k, v in lastocc.items():
            freq = cnt[k]
            # need to read one less
            if k == cchar:
                freq -= 1

            # already have enough of this character
            if freq == 0:
                continue

            # need more than what we have to build window
            if freq > len(lastocc[k]):
                return None
            
            mpos = min(mpos, lastocc[k][len(lastocc[k]) - freq])

        return mpos




s = Solution()
print(s.minWindow("sdadsadhsladjlasd", "sadd"))

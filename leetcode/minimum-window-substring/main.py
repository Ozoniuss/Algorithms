from collections import Counter, defaultdict
import math

class Solution:


    def minWindow(self, s: str, t: str) -> str:

        if len(t) == 0:
            return ""

        if len(t) == 1:
            if t in s:
                return t
            else:
                return ""
            
        origs = s
        cnt = dict(Counter(t))
        needchars = len(cnt)

        s = list((idx, c) for idx, c in enumerate(s) if c in cnt)
        actualctt = defaultdict(int)
        gotchars = 0

        l, r = 0,0
        minstr, minstrlen = "", math.inf
        while r < len(s):
            rpos, rchar = s[r][0], s[r][1]
            actualctt[rchar] += 1
            if actualctt[rchar] == cnt[rchar]:
                gotchars += 1
            if gotchars >= needchars:
                while l <= r:
                    lpos, lchar = s[l][0], s[l][1]
                    if actualctt[lchar] - 1 >= cnt[lchar]:
                        actualctt[lchar] -= 1
                        l += 1
                    else:
                        if (rpos-lpos + 1) < minstrlen:
                            minstrlen = (rpos-lpos+1)
                            minstr = origs[lpos:rpos+1]
                        break
            r += 1
        return minstr




    def minWindowv2(self, s: str, t: str) -> str:

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
print(s.minWindow("ababab", "aaabbb"))
print(s.minWindow("abab", "aaabbb"))
print(s.minWindow("ADOBECODEBANC", "ABC"))
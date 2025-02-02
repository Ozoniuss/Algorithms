class Solution:
    out = []
    def longestValidParentheses(self, s: str) -> int:
        self.out = []
        stk = []
        for i, c in enumerate(s):
            if c == '(':
                stk.append(i)
            elif c == ')':
                if len(stk) == 0:
                    continue
                else:
                    pos = stk.pop()
                    self.appendOut((pos, i+1))

        maxlen = 0
        for el in self.out:
            maxlen = max(maxlen, el[1] - el[0])
        return maxlen
    
    def appendOut(self, tup):
        if len(self.out) == 0:
            self.out.append(tup)
            return 
        ptup = self.out[-1]
        if tup[0] < ptup[1]:
            self.out = self.out[:-1] # remove last element
            self.appendOut(tup)
        elif tup[0] == ptup[1]:
            self.out[-1] = (ptup[0], tup[1]) # extend
            return
        else:
            self.out.append(tup)
            return

    

s = Solution()
print(s.longestValidParentheses(""))
# print(s.longestValidParentheses("()(()())"))
# print(s.longestValidParentheses("(()"))
# print(s.longestValidParentheses("(()()"))
# print(s.longestValidParentheses("(()(()"))
# print(s.longestValidParentheses("(()(()))"))
# print(s.longestValidParentheses("(()(()()))"))
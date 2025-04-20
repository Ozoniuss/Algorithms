from typing import List
from collections import deque

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        s = deque()
        for t in tokens:
            if t in "+-/*":
                assert(len(s) >= 2)
                op2 = s.pop()
                op1 = s.pop()
                r = None
                if t == "+":
                    r = op1 + op2
                elif t == "-":
                    r = op1 - op2
                elif t == "/":
                    print(op1, op2, op1//op2)
                    r = int(op1 / op2)
                elif t == "*":
                    r = op1 * op2
                s.append(r)
            else:
                s.append(int(t))
        return s[0]
    

s = Solution()
# print(s.evalRPN(["2","1","+","3","*"]))
# print(s.evalRPN(["4","13","5","/","+"]))
print(s.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]))
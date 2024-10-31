import math
from typing import List


class Solution:
    def div(self, a, b):
        if a / b > 0:
            return a // b
        return int(math.ceil(a / b))

    def evalRPN(self, tokens: List[str]) -> int:
        arguments = []
        for t in tokens:
            if t in "+-/*":
                arg1 = int(arguments[-2])
                arg2 = int(arguments[-1])
                arguments = arguments[:-2]

                val = 0
                if t == "+":
                    val = arg1 + arg2
                if t == "-":
                    val = arg1 - arg2
                if t == "*":
                    val = arg1 * arg2
                if t == "/":
                    val = self.div(arg1, arg2)
                arguments.append(val)
            else:
                arguments.append(int(t))

            print(arguments)
        return arguments[0]


s = Solution()
print(
    s.evalRPN(
        tokens=["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
    )
)

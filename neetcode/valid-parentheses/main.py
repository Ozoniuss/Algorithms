class Solution:
    def isValid(self, s: str) -> bool:
        pstack = []
        matches = {"{": "}", "[": "]", "(": ")"}
        for letter in s:
            if letter in ["{", "[", "("]:
                pstack.append(letter)
            elif letter in ["}", "]", ")"]:
                if len(pstack) == 0 or letter != matches[pstack[-1]]:
                    return False
                pstack.pop()
        return len(pstack) == 0

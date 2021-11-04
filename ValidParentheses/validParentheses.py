class Solution:
    def isValid(self, s: str) -> bool:

        stack = [] # stores the order of which types of brackets are open
        # basically, if we open a type of bracket, and we open another type
        # inside that bracket, we must close the second bracket before opening a new one.

        # it will contain only open brackets

        matches = {'(':')',
                   '{':'}',
                   '[':']'}

        for bracket in s:
            if bracket in '({[': # if bracket is open, we add it to the stack
                stack.append(bracket)
            elif bracket in ')}]': # if bracket is closed, we must remove the same type of bracket from the stack
                if len(stack) == 0: # if the stack is opened there was no opened bracket
                    return False
                last_bracket = stack[-1]
                if matches[last_bracket] == bracket: #
                    stack.pop()
                else:
                    return False

        return len(stack) == 0

s = Solution()
print(s.isValid('{[]}'))
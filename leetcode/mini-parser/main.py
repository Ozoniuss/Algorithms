# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
class NestedInteger:
    def __init__(self, value=None):
        """
        If value is not specified, initializes an empty list.
        Otherwise initializes a single integer equal to value.
        """

    def isInteger(self):
        """
        @return True if this NestedInteger holds a single integer, rather than a nested list.
        :rtype bool
        """

    def add(self, elem):
        """
        Set this NestedInteger to hold a nested list and adds a nested integer elem to it.
        :rtype void
        """

    def setInteger(self, value):
        """
        Set this NestedInteger to hold a single integer equal to value.
        :rtype void
        """

    def getInteger(self):
        """
        @return the single integer that this NestedInteger holds, if it holds a single integer
        Return None if this NestedInteger holds a nested list
        :rtype int
        """

    def getList(self):
        """
        @return the nested list that this NestedInteger holds, if it holds a nested list
        Return None if this NestedInteger holds a single integer
        :rtype List[NestedInteger]
        """


class Solution:
    def deserialize(self, s: str):
        if s[0] not in "[]":
            return NestedInteger(value=int(s))
        depth = 0
        integers = []
        start = 0
        inner = s[1:-1]
        for idx, c in enumerate(inner):
            if c == "[":
                depth += 1
            elif c == "]":
                depth -= 1

            if c == "," and depth == 0:
                print(inner[start:idx])
                integers.append(self.deserialize(inner[start:idx]))
                start = idx + 1
            if idx == len(inner) - 1:
                print(inner[start : idx + 1])
                integers.append(self.deserialize(inner[start : idx + 1]))

        i = NestedInteger()
        for itm in integers:
            i.add(itm)
        return i


s = Solution()
s.deserialize("[1,[2,[3]]]")

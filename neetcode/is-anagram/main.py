from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str):
        counter = Counter(s)
        counterDict = dict(counter)

        for element in t:
            if element in counterDict:
                counterDict[element] -= 1
                if counterDict[element] == 0:
                    del counterDict[element]
            else:
                return False

        if len(counterDict) != 0:
            return False

        return True

    def isAnagramFrequency(self, s: str, t: str):
        freq = [0] * 26
        for element in s:
            freq[element - "a"] += 1

        for element in t:
            freq[element - "a"] -= 1

        return not any(freq)


print(any([0, 0, 1]))
s = Solution()
print(s.isAnagram("xx", "x"))

from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for s in strs:
            sortedString = str(sorted(s))
            if sortedString not in groups:
                groups[sortedString] = [s]
            else:
                groups[sortedString].append(s)

        return list(groups.values())


s = Solution()

print(s.groupAnagrams(["act", "pots", "tops", "cat", "stop", "hat"]))

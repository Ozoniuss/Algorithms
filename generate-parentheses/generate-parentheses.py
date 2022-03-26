from typing import List

class Solution:

    #naive approach using a hashset
    def generateParenthesis(self, n: int) -> List[str]:
        allResults = set()
        result = [""]
        for i in range(1,n+1):
            new_results = set()
            for comb1 in allResults:
                for comb2 in allResults:
                    if len(comb1 + comb2) == 2*i:
                        new_results.add(comb1 + comb2)
                        new_results.add(comb2 + comb1)
            allResults = allResults.union(new_results)
            for validCombination in result:
                new_results.add('(' + validCombination + ')')
                allResults.add('(' + validCombination + ')')


            result = list(new_results)

        return result

s = Solution()
for idx, vp in enumerate(s.generateParenthesis(4)):
    print(vp)
    print(idx)

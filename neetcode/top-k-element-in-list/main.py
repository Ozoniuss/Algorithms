from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = [(0, i) for i in range(2001)]
        for num in nums:
            fq, pos = freq[num + 1000]
            freq[num + 1000] = (fq + 1, pos)

        out = []
        for elements in sorted(freq, key=lambda x: x[0], reverse=True):
            out.append(elements[1] - 1000)
            if len(out) == k:
                break

        return out


s = Solution()
print(s.topKFrequent([1, 2, 2, 3, 3, 3], 2))

from typing import List


class Solution:

    def getLength(self, num: int, sequences: dict):
        return abs(num - sequences[num]) + 1

    def longestConsecutive(self, nums: List[int]) -> int:
        # basically keep a pointer to the bound of the longest sequence
        # it belongs to

        sequences = {}
        longest = 0

        for num in nums:

            # this number already belongs to some sequence
            if num in sequences:
                continue

            prev, next = num - 1, num + 1
            if prev not in sequences and next not in sequences:
                # num is only in a consecutive sequence containing itself
                sequences[num] = num
                longest = max(self.getLength(num, sequences), longest)
            elif prev in sequences and next not in sequences:
                # there is only a "left" consecutive sequence that num could
                # be part of. extend that sequence

                # point num to the beginning of the sequence
                sequences[num] = sequences[prev]

                # point the beginning of the sequence to num
                sequences[sequences[num]] = num
                longest = max(self.getLength(num, sequences), longest)
            elif prev not in sequences and next in sequences:
                # do the same for the right sequence
                sequences[num] = sequences[next]
                sequences[sequences[num]] = num
                longest = max(self.getLength(num, sequences), longest)
            else:
                # this is the only number missing from the two sequences
                # merge the two sequences together

                # doesn't matter, just have num there in case it repeats
                # be sure to use a number that doesn't alter sequences
                sequences[num] = num

                mx = sequences[next]
                mn = sequences[prev]

                sequences[mx] = mn
                sequences[mn] = mx
                longest = max(self.getLength(sequences[mn], sequences), longest)

        return longest


s = Solution()
print(s.longestConsecutive([100, 4, 200, 1, 3, 2]))
print(s.longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))
print(s.longestConsecutive([0]))

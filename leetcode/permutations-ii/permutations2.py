from typing import List

class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        # naive approach that stores all possible permitations in a set
        d = {i:nums[i] for i in range(len(nums))}
        allPermutationsSoFar = set()

        keys_permutations = [[]]

        for _ in range(len(nums)):
            new_permutations = []
            for p in keys_permutations:
                preset_elements = set(p)
                for key2 in range(len(nums)):
                    if key2 not in preset_elements:
                        new_permutations.append(p+[key2])
            keys_permutations = new_permutations

        final_permutations = []

        for key_permutation in keys_permutations:
            actual_permutation = [d[k] for k in key_permutation]
            strint_actual_permutation = ''.join([str(i) for i in actual_permutation])
            if strint_actual_permutation not in allPermutationsSoFar:
                allPermutationsSoFar.add(strint_actual_permutation)
                final_permutations.append(actual_permutation)

        return final_permutations

    def permuteUniqueRecursive(self, nums: List[int]) -> List[List[int]]:
        # this is a more clever approach using recursion
        if len(nums) == 0:
            return [[]]

        nums.sort()
        # basically we are going to sort the numbers, and only go through the unique elements
        # the reasoning behind would be that at each step we add some number to the permutations,
        # but when we add the same number we generate the same permutations

        copy_nums = nums.copy() # copy to avoid ruining the set when removing
        result = []

        for number in set(nums): # go only through the unique numbers
            copy_nums.remove(number)
            permutations = self.permuteUniqueRecursive(copy_nums)
            copy_nums.append(number)

            for p in permutations:# add the permutations of the set not containing that number
                result.append(p + [number])

        return result


s = Solution()
print(s.permuteUniqueRecursive(nums = [1,1,2]))

# a = [1,1,1,2,3]
# a.remove(1)
# print(a)
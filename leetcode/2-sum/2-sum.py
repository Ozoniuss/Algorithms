from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # note that we can even do this in a single attempt, but it's probably not worth the imporvement

        numbers_dict = {} # we will use a dictionary to store the numbers as keys, and their indexes as values (list)
        for idx, n in enumerate(nums):
            if n not in numbers_dict:
                numbers_dict[n] = [idx]
            else:
                # if the number is already in the hashmap, it's present with a different index as well
                numbers_dict[n].append(idx)

        for number in nums:
            if target - number in numbers_dict: # basically check if the sum is target
                if target - number == number:
                    # if the numbers are equal, then we must have exactly two of them to ensure that we didn't add twice
                    if len(numbers_dict[number]) == 1:
                        continue
                    else:
                        return numbers_dict[number] # since the solution is unique
                else:
                    return [numbers_dict[number][0], numbers_dict[target - number][0]]


s = Solution()
print(s.twoSum(nums = [3,2,4], target = 6))
from typing import List

# This one is ez af

class Solution:

    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left_index = 0
        right_index = len(numbers) - 1
        #basically, if the sum is greater, we have to decrease the right index to lower the sum, and if the
        #sum is smaller, we have to increase the left index to increase the sum
        # O(n) complexity

        while True: # there is always a solution so no need to stop
            if numbers[left_index] + numbers[right_index] > target:
                right_index -= 1
            elif numbers[left_index] + numbers[right_index] < target:
                left_index += 1
            else:
                return [left_index+1, right_index+1]

s = Solution()
print(s.twoSum(numbers = [2,7,11,15], target = 9))
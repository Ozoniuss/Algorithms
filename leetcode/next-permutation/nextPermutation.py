from typing import List

# nice solution, faster than 95% on speed and 99% on memory on leetcode
# complexity: between O(n) and O(2n)

class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        # define two auxiliary functions. We'll use this first one to reverse a sorted list in O(n/2)
        def swap(index1, index2):
            # swap the positions of two numbers
            aux = nums[index1]
            nums[index1] = nums[index2]
            nums[index2] = aux

        # assume that the elements from index are sorted in descending order
        def re_sort(index):
            # re sort starting from index
            lastIndex = len(nums) - 1
            for i in range((lastIndex - index + 1) // 2):
                swap(index + i, lastIndex - i)


        # the idea is to find the last element that is smaller than the succeeding element.
        # we know that all numbers will be sorted in descending order after this element
        last_smaller_index = -1
        for i in range(len(nums) - 1):
            if nums[i] < nums[i+1]:
                last_smaller_index = i # this is the index of the last element smaller than the succeeding,
                # call the element e0

        # if we didn't find such element, the elements were sorted descendingly, and hence
        # we have to return the first permutation, in ascending order

        if last_smaller_index == -1:
            re_sort(0) # which means resorting from 0 in place

        else:
            # basically, the idea here is to find in the remaining elements
            # the last element that is greater than the one we found earlier, e1

            # we are going to swap that one with e0
            # because the elements after e0 are descending, e0 will be replaced with e1 for the alphabetic order
            # the elements after e1 now will be in descending order, so just resort those and we're done

            currentLastSmallerElement = nums[last_smaller_index]
            needsToSwapWithIndex = None
            for i in range(last_smaller_index + 1, len(nums)):
                # find e1 we talked about earlier
                if nums[i] <= currentLastSmallerElement:
                    break
                needsToSwapWithIndex = i

            swap(last_smaller_index, needsToSwapWithIndex) # swap e0 with e1
            re_sort(last_smaller_index+1) # sort elements after e1




s = Solution()
nums = [5,1,1]
s.nextPermutation(nums)
print(nums)


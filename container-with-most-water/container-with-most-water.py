from typing import List

class Solution:
    def maxAreaBruteForce(self, height: List[int]) -> int:

        # brute force approach, check all possible combinations of walls
        # complexity: O(h^2)

        max_area = 0

        for i in range(0, len(height) - 1):
            for j in range(i+1, len(height)):
                area = min(height[i], height[j]) * (abs(i-j))
                if area > max_area:
                    max_area = area
        return max_area

    def maxAreaLeftRight(self, height: List[int]) -> int:

        # clever approach, start with a pointer from left and another from right
        # the idea behind this one is to compare the size of the left wall and the size of the right wall
        # say the left one is smaller, then we move the pointer to the right
        # unless we encounter a wall that is at least as tall as the left one, the new area will always be smaller,
        # since the smaller wall is actually the height
        # if we encounter a taller wall, it will be our new left
        # repeat
        # total complexity: O(h)


        left = 0 # leftmost pointer
        right = len(height) - 1 # rightmost pointer

        max_area = (len(height) - 1) * min(height[0], height[-1]) # initial area

        while True:
            if height[left] > height[right]: # go from the smaller one to the bigger one
                current_position = right
                while (height[current_position] <= height[right]):
                    current_position -= 1 # go one to the left until you find a taller one
                    if current_position == left: # if reaches the left, just break
                        return max_area

                # if did not reach left, set the new location of right
                right = current_position
                area = (abs(left - right)) * min(height[left], height[right])
                if area > max_area:
                    max_area = area

            else:
                current_position = left
                while (height[current_position] <= height[left]):
                    current_position += 1  # go one to the right until you find a bigger one
                    if current_position == right:  # if reached the right, just break
                        return max_area

                # if did not reach right, set the new location of left
                left = current_position
                area = (abs(left - right)) * min(height[left], height[right])
                if area > max_area:
                    max_area = area


s = Solution()
print(s.maxAreaLeftRight(height = [1,2,1]))

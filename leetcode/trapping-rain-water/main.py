import enum
from typing import List

class Solution:
    
    def trap(self, height: List[int]) -> int:

        new_heights = height.copy()
        height_adjusted = [-1] + height + [-1]
        local_maximums = []

        last_checked = -1
        for i in range(1, len(height_adjusted) - 1):
            if i < last_checked:
                continue
            l, pl = height_adjusted[i-1], i-1
            m, pm = height_adjusted[i], i
            r, pr = height_adjusted[i+1], i+1
            if height_adjusted[i] == height_adjusted[i+1]:
                for j in range(i+1, len(height_adjusted)):
                    # guaranteed to happen because of padding
                    if height_adjusted[i] != height_adjusted[j]:
                        r, pr = height_adjusted[j], j
                        break

            last_checked = pr
            if l<m>r:
                if len(local_maximums) == 0:
                    local_maximums.append((m,i-1))

                else:
                    
                    largest_right = -1
                    largest_right_pos = -1
                    # find larger, if necessary
                    for ll in local_maximums:
                        mm, pp = ll
                        if mm >= largest_right:
                            largest_right = mm
                            largest_right_pos = pp
                    
                    local_maximums.append((m,i-1))

                    # something will block
                    if m < largest_right:
                        for j in range(i-2, -1, -1):
                            if new_heights[j] >= m:
                                break
                            if new_heights[j] < m:
                                new_heights[j] = m # fill

                    else:
                        for j in range(i-2, largest_right_pos-1, -1):
                            if new_heights[j] < largest_right:
                                new_heights[j] = largest_right # fill

        # print(local_maximums)
        # print(new_heights)
        return sum(new_heights) - sum(height)

    # figured after hints
    def trap_v2(self, height: List[int]) -> int:


        maxleft = [0] * len(height)
        maxright = [0] * len(height)
        m = 0
        for idx, h in enumerate(height):
            maxleft[idx] = m
            if h > m:
                m = h
        m = 0
        for idx, h in enumerate(reversed(height)):
            maxright[-idx-1] = m
            if h > m:
                m = h
        # print(maxleft, maxright)
        s = 0
        for idx, h in enumerate(height):
            # print(idx, h)
            s += max(0, min(maxleft[idx],maxright[idx]) - h)


        return s

    # pretty clever solution I didn't come up with
    def trap_v3(self, height: List[int]) -> int:

        left, right, idx = 0, len(height)-1, 0
        maxleft = 0
        maxright = 0
        s = 0
        while left != right:
            if maxleft <= maxright:
                s += max(0, min(maxleft, maxright) - height[left])
                left += 1
                maxleft = max(maxleft, height[left-1])
            else:
                s += max(0, min(maxleft, maxright) - height[right])
                right -= 1
                maxright = max(maxright, height[right+1])
        s += max(0, min(maxleft, maxright) - height[left])
        
        return s

print(list((i, j) for (i, j) in enumerate(reversed([1,2,3]))))

s = Solution()
# print(s.trap([2,0,3,2,2,0,1,0,4,0,3,1])) #  2 2 3 3 3 3 3 3 4 3 3 1
# print(s.trap([5,0,3,0,2,0,2,0,3])) # 5 3 3 3 3 3 3 3 3
# print(s.trap([0,1,0,2,1,0,1,3,2,1,2,1]))  
# print(s.trap_v3([5,5,4,7,8,2,6,9,4,5]))  # 5 5 5 7 8 8 8 9 5 5 
print(s.trap_v3([5,5,1,7,1,1,5,2,7,6])) # 5 5 5 7 7 7 7 7 7 6 
# print(s.trap([7,7,5,5,7,7])) # 5 5 5 7 7 7 7 7 7 6 
print(s.trap_v3([4,2,0,3,2,5])) # 5 5 5 7 7 7 7 7 7 6 
  
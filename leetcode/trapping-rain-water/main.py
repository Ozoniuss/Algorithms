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
            # print("local extremum", pl, pm, pr, height_adjusted)
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
        

s = Solution()
# print(s.trap([2,0,3,2,2,0,1,0,4,0,3,1])) #  2 2 3 3 3 3 3 3 4 3 3 1
# print(s.trap([5,0,3,0,2,0,2,0,3])) # 5 3 3 3 3 3 3 3 3
# print(s.trap([0,1,0,2,1,0,1,3,2,1,2,1]))  
print(s.trap([5,5,4,7,8,2,6,9,4,5]))  # 5 5 5 7 8 8 8 9 5 5 
# print(s.trap([5,5,1,7,1,1,5,2,7,6])) # 5 5 5 7 7 7 7 7 7 6 
# print(s.trap([7,7,5,5,7,7])) # 5 5 5 7 7 7 7 7 7 6 
  
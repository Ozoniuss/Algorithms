from typing import List

class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key = lambda x: (x[0], x[1]))
        kept = []
        # print(intervals)
        for i in intervals:
            if len(kept) == 0:
                kept.append(i.copy())
                continue

            [k1, k2] = kept[-1]
            [i1, i2] =  i

            # do not keep no questions asked
            if k1 == i1:
                continue 

            # i1 > k1

            if i1 >= k2:
                kept.append([i1, i2])
                continue

            # k1 < i1 < k2
            if i2 <= k2:
                kept[-1] = [i1, i2]
                continue
 
        return len(intervals) - len(kept)

    def eraseOverlapIntervalsWithLastOnly(self, intervals: List[List[int]]) -> int:
        intervals.sort(key = lambda x: (x[0], x[1]))
        last = None
        c = 0

        for i in intervals:
            if last is None:
                last = i.copy()
                c += 1
                continue

            [k1, k2] = last
            [i1, i2] =  i

            # do not keep no questions asked
            if k1 == i1:
                continue 

            if i1 >= k2:
                last = [i1, i2]
                c += 1

            # k1 < i1 < k2
            if i2 <= k2:
                last = [i1, i2]
                continue
 
        return len(intervals) - c

s = Solution()
print(s.eraseOverlapIntervals([[1,2],[2,3],[3,4],[1,3]]))
print(s.eraseOverlapIntervals([[1,2],[1,2],[1,2]]))
print(s.eraseOverlapIntervals([[1,2],[2,3]]))
print(s.eraseOverlapIntervals([[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]))

print(s.eraseOverlapIntervalsWithLastOnly([[1,2],[2,3],[3,4],[1,3]]))
print(s.eraseOverlapIntervalsWithLastOnly([[1,2],[1,2],[1,2]]))
print(s.eraseOverlapIntervalsWithLastOnly([[1,2],[2,3]]))
print(s.eraseOverlapIntervalsWithLastOnly([[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]))


# l = [[3, 6], [3, 5], [1,7]]
# l.sort(key=lambda x: (x[0], x[1]))
# print(l)
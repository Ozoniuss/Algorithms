from collections import deque


class Solution(object):
    def highestPeak(self, isWater):
        """
        :type isWater: List[List[int]]
        :rtype: List[List[int]]
        """
        q = deque()
        LL = len(isWater)
        CC = len(isWater[0])
        elevation = [[-1 for j in range(CC)] for i in range(LL)]
        for i in range(LL):
            for j in range(CC):
                if isWater[i][j] == 1:
                    q.append((i, j, 0))
                    elevation[i][j] = 0

        c = 0
        while len(q) != 0:
            top = q.popleft()

            i, j, ln = top
            c += 1
            print(c, top)

            # if ln != 0:
            #     elevation[i][j] = ln

            ns = [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]
            for n in ns:
                if not (0 <= n[0] < LL and 0 <= n[1] < CC):
                    continue
                if elevation[n[0]][n[1]] != -1:
                    continue
                elevation[n[0]][n[1]] = ln + 1
                q.append((n[0], n[1], ln + 1))

        return elevation


s = Solution()
print(s.highestPeak([[0, 0, 0, 1], [0, 0, 1, 0], [1, 0, 0, 0]]))
print(s.highestPeak([[0, 0, 1], [1, 0, 0], [0, 0, 0]]))
print(s.highestPeak([[0, 1], [0, 1]]))

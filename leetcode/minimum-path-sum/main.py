from typing import List


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        L, LL = 0, len(grid)
        C, CC = 0, len(grid[0])

        dp = [[0 for j in range(CC + 1)] for i in range(LL + 1)]
        for c in range(CC):
            dp[1][c + 1] = grid[0][c] + dp[1][c]
        for l in range(LL):
            dp[l + 1][1] = grid[l][0] + dp[l][1]
        print(dp)
        for i in range(1, LL):
            for j in range(1, CC):
                k = grid[i][j]
                dp[i + 1][j + 1] = min(k + dp[i][j + 1], k + dp[i + 1][j])
        print(dp)
        return dp[LL][CC]

    def minPathSumBfs(self, grid: List[List[int]]) -> int:
        L, LL = 0, len(grid)
        C, CC = 0, len(grid[0])

        dp = [[0 for j in range(CC)] for i in range(LL)]

        parents = {}
        round = 0
        maxround = LL + CC

        while round < maxround:
            current = []
            for i in range(round + 1):
                j = round - i
                if 0 <= i < LL and 0 <= j < CC:
                    current.append((i, j))

            for c in current:
                i, j = c
                k = grid[i][j]
                mmax = 999999999999999
                n1 = dp[i - 1][j] if 0 <= i - 1 < LL and 0 <= j < CC else mmax
                n2 = dp[i][j - 1] if 0 <= i < LL and 0 <= j - 1 < CC else mmax
                if n1 == n2 == mmax:
                    dp[i][j] = k
                else:
                    dp[i][j] = min(k + n1, k + n2)
                # if n1 == n2 == max:
                #     parents[(i, j)] = None
                # elif n1 < n2:
                #     parents[(i, j)] = (i - 1, j)
                # else:
                #     parents[(i, j)] = (i, j - 1)

            round += 1

        return dp[LL - 1][CC - 1]
        # path = []
        # current = (LL - 1, CC - 1)
        # while current != (0, 0):
        #     path.append(current)
        #     current = parents[current]
        # path.append((0, 0))
        # return list(reversed(path))


s = Solution()
print(s.minPathSum([[1, 3, 1], [1, 5, 1], [4, 2, 1]]))
# print(s.minPathSum([[1, 2, 3], [4, 5, 6]]))

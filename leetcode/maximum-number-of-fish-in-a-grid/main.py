from typing import List


class Solution:

    def explore(self, round, current, grid, LL, CC, visited):
        if current in round:
            return 0
        if grid[current[0]][current[1]] == 0:
            return 0
        round[current] = grid[current[0]][current[1]]
        visited.add(current)
        i, j = current
        ns = [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]
        t = grid[current[0]][current[1]]
        for n in ns:
            if 0 <= n[0] < LL and 0 <= n[1] < CC and n not in round:
                t += self.explore(round, n, grid, LL, CC, visited)
        return t

    def findMaxFish(self, grid: List[List[int]]) -> int:
        LL = len(grid)
        CC = len(grid[0])

        visited = set()
        mtotal = 0

        for i in range(LL):
            for j in range(CC):
                if (i, j) in visited:
                    continue
                if grid[i][j] == 0:
                    continue
                round = {}
                s = self.explore(round, (i, j), grid, LL, CC, visited)
                if s > mtotal:
                    mtotal = s

        return mtotal


# grid = [[0, 2, 1, 0], [4, 0, 0, 3], [1, 0, 0, 4], [0, 3, 2, 0]]
grid = [[0]]
s = Solution()
print(s.findMaxFish(grid))

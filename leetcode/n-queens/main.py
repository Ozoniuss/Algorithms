from typing import List

class Solution:

    def __init__(self):
        self.solutions = []

    @staticmethod
    def deepcopy(mask: List[List[int]]):
        n = len(mask)
        return [[e for e in mask[i]] for i in range(n)]

    @staticmethod
    def toSolution(n, placed: List[int]):
        out = []
        for col in placed:
            s = ["." for i in range(n)]
            s[col] = 'Q'
            out.append("".join(s))
        return out
        

    def solveNQueens(self, n: int) -> List[List[str]]:
        mask = [[0 for i in range(n)] for i in range(n)]
        placed = [-1 for i in range(n)]
        for i in range(n):
            # place queen on column
            m = self.deepcopy(mask)
            placed[0] = i
            self.fill_mask(n, m, 0, i)
            self.solve_rec(n, 1, placed, m)
            placed[0] = -1

        return self.solutions

    def solve_rec(self, n, npos, placed, mask):

        if npos == n:
            self.solutions.append(self.toSolution(n,placed))
            return
        
        for col in range(n):
            if mask[npos][col] == 1:
                continue

            m = self.deepcopy(mask)
            placed[npos] = col
            self.fill_mask(n, m, npos, col)
            self.solve_rec(n,npos+1, placed,self.deepcopy(m))
            placed[npos] = -1






    def fill_mask(self, n, mask, x, y):
        for i in range(n):
            mask[x][i] = 1
        for i in range(n):
            mask[i][y] = 1

        i = 0
        while True:
            changed = False
            if x-i >= 0 and y-i >= 0:
                mask[x-i][y-i] = 1
                changed = True
            if x+i < n and y+i <n:
                mask[x+i][y+i] = 1
                changed = True
            if not changed:
                break
            i += 1

        i = 0
        while True:
            changed = False
            if x+i < n and y-i >= 0:
                mask[x+i][y-i] = 1
                changed = True
            if x-i >= 0 and y+i < n:
                mask[x-i][y+i] = 1
                changed = True
            if not changed:
                break
            i += 1

   

s = Solution()
print(s.toSolution(4, [1,0,3,2]))
print(s.solveNQueens(9))

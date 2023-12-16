from typing import List


class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        L = len(mat)
        diagonals = [mat[i][i] + mat[i][L - i - 1] for i in range(L)]
        minus = mat[L // 2][L // 2] if L % 2 != 0 else 0
        return sum(diagonals) - minus

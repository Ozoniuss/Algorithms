from typing import List

matrix = [[0 for i in range(3)] for j in range(4)]
print(matrix)


class Solution:
    def spiralMatrixIII(
        self, rows: int, cols: int, rStart: int, cStart: int
    ) -> List[List[int]]:
        boundaries = [rStart, rStart + 1, cStart, cStart + 2]
        current = 1
        total = rows * cols

        matrix = [[0 for j in range(cols)] for i in range(rows)]
        out = []

        while True:

            [L, LL, C, CC] = boundaries
            i = L
            for j in range(C, CC):
                if (0 <= i < rows) and (0 <= j < cols) and matrix[i][j] == 0:
                    matrix[i][j] = current
                    out.append([i, j])
                    current += 1
                    if current == total + 1:
                        return out
            LL += 1
            boundaries = [L, LL, C, CC]

            [L, LL, C, CC] = boundaries
            j = CC - 1
            for i in range(L, LL):
                if (0 <= i < rows) and (0 <= j < cols) and matrix[i][j] == 0:
                    matrix[i][j] = current
                    out.append([i, j])
                    current += 1
                    if current == total + 1:
                        return out
            C -= 1
            boundaries = [L, LL, C, CC]

            [L, LL, C, CC] = boundaries
            i = LL - 1
            for j in range(CC - 1, C - 1, -1):
                if (0 <= i < rows) and (0 <= j < cols) and matrix[i][j] == 0:
                    matrix[i][j] = current
                    out.append([i, j])
                    current += 1
                    if current == total + 1:
                        return out
            L -= 1
            boundaries = [L, LL, C, CC]

            [L, LL, C, CC] = boundaries
            j = C
            for i in range(LL - 1, L - 1, -1):
                if (0 <= i < rows) and (0 <= j < cols) and matrix[i][j] == 0:
                    matrix[i][j] = current
                    out.append([i, j])
                    current += 1
                    if current == total + 1:
                        return out

            CC += 1
            boundaries = [L, LL, C, CC]


s = Solution()

print(s.spiralMatrixIII(1, 4, 0, 0))

s = Solution()
print(s.spiralMatrixIII(5, 6, 1, 4))

from typing import List


class Solution:

    def hasDuplicates(self, elements: List[str]) -> bool:
        existing = set()
        for element in elements:
            if element == ".":
                continue
            elif element not in "123456789":
                return False
            elif element in existing:
                return True
            else:
                existing.add(element)
        return False

    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # check all rows
        for row in range(9):
            if self.hasDuplicates(board[row]):
                return False

        for column in range(9):
            elements = [board[row][column] for row in range(9)]
            if self.hasDuplicates(elements):
                return False

        for row in range(0, 9, 3):
            for col in range(0, 9, 3):
                elements = []
                elements.extend(board[row][col : col + 3])
                elements.extend(board[row + 1][col : col + 3])
                elements.extend(board[row + 2][col : col + 3])

                if self.hasDuplicates(elements):
                    return False

        return True


board = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", ".", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

board2 = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "1", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", ".", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

s = Solution()
print(s.isValidSudoku(board))
print(s.isValidSudoku(board2))

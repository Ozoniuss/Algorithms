from typing import List

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # this is quite ez
        # you have to check all possibilities
        # and there aren't even that many to begin with

        # go through lines
        for line in board:
            elements = set()
            for element in line:
                if element!='.':
                    if element in elements:
                        return False
                    elements.add(element)

        # go through columns
        for column in range(9):
            elements = set()
            for line in range(9):
                element = board[line][column]
                if element != '.':
                    if element in elements:
                        return False
                    elements.add(element)

        # 9 squares
        left = (0,1,2)
        middle = (3,4,5)
        right = (6,7,8)

        squares = [(left, left),
                   (left, middle),
                   (left, right),
                   (middle, left),
                   (middle, middle),
                   (middle, right),
                   (right, left),
                   (right, middle),
                   (right, right),]

        for square in squares:
            elements = set()
            for line in square[0]:
                for column in square[1]:
                    element = board[line][column]
                    if element != '.':
                        if element in elements:
                            return False
                        elements.add(element)

        return True

s = Solution()

print(s.isValidSudoku(board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]))
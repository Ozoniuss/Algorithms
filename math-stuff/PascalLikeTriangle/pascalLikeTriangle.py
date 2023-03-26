# #have the idea, will implement later
#
#
# from typing import List
#
#
# # we assume that there are n(n+1)/2 numbers, which means n rows
# def sortNumbersInLine(line: List[int]):
#     """
#     This function does the sorting in the following way:
#     the list is sorted increasingly, so we move through the list and
#     add the first number to the left, the second to the right, and so on
#     In this way the numbers should also match under the pyramid, since the sum of
#     the first two is going to be the greatest, followed by the sum of the last two,
#     followed by the sum of the second and third, etc which is the same ordering as the row above
#     """
#     leftOrdering = []
#     rightOrdering = []
#     print("a" + str(line))
#     for idx, num in enumerate(line):
#         if idx % 2 == 0:  # these go on  the left
#             leftOrdering.append(num)
#         else:
#             rightOrdering.insert(0, num)
#     leftOrdering.extend(rightOrdering)
#     return leftOrdering
#
#
# def findNumbersOneachRow(numbers: List[int], n: int):
#     """This function returns a dict with the keys being the row numbers, starting from one at the top,
#     and the values being a list of numbers, sorted in the correct order"""
#     # n is the length of the list
#
#     numbers.sort(reverse=False)
#     print(numbers)
#     pyramidRows = {}
#
#     for i in range(1, n + 1):
#         pyramidRows[i] = []
#         for j in range(1, i+1):
#             pyramidRows[i].insert(0, numbers.pop())
#
#         pyramidRows[i] = sortNumbersInLine(pyramidRows[i])
#
#     return pyramidRows
#
#
# # d = (findNumbersOneachRow([1, 1, 2, 3, 3, 4, 5, 8, 9, 17], 4))
# d = (findNumbersOneachRow([10,5,6,1,4,3], 3))
# for k in d:
#     print(d[k])

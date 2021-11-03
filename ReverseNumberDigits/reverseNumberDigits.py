class Solution:
    def reverse(self, x: int) -> int:
        s = str(x)
        reversed = ''
        for letter in s:
            reversed = letter + reversed

        if reversed[-1] == '-':
            reversed = "-" + reversed[:-1]

        number = int(reversed)
        if (-2 ** 31 <= number <= 2 ** 31 - 1):
            return number

        return 0



s = Solution()
print(s.reverse(-120))

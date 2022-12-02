from typing import List


# A classic bfs iterative solution and a recursive solution

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if len(digits) == 0:
            return []

        result = [""]

        combinations = {"2":"abc",
                        "3":"def",
                        "4":"ghi",
                        "5":"jkl",
                        "6":"mno",
                        "7":"pqrs",
                        "8":"tuv",
                        "9":"wxyz"}

        for digit in digits:
            new_combinations = []
            corresponding_letters = combinations[digit]

            for p in result: # go through all pre-existing combinations
                for letter in corresponding_letters: # add each possible letter to each permutation
                    new_combinations.append(p + letter)

            result = new_combinations

        return result

    def letterCombinationsRecursive(self, digits: str) -> List[str]:

        if len(digits) == 0:
            return []


        result = []

        combinations = {"2":"abc",
                        "3":"def",
                        "4":"ghi",
                        "5":"jkl",
                        "6":"mno",
                        "7":"pqrs",
                        "8":"tuv",
                        "9":"wxyz"}

        if len(digits) == 1:
            return list(combinations[digits[0]])

        firstDigit = digits[0]
        lastDigits = digits[1:]
        print(digits)
        print(lastDigits)

        laterCombinations = self.letterCombinationsRecursive(lastDigits)
        for letter in combinations[firstDigit]:
            for c in laterCombinations:
                print(c)
                result.append(letter + c)

        return result


s = Solution()
print(s.letterCombinationsRecursive("24"))


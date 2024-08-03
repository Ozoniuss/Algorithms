from typing import List

# Double each letter in the word and add the "01" separator between two words.
# This ensures that there will always be two equal consecutive letters, unless
# we transition between two strings.


class Solution:

    def encode(self, strs: List[str]) -> str:
        sep = "01"
        out = ""
        for word in strs:
            for c in word:
                out += c * 2
            out += sep
        return out

    def decode(self, s: str) -> List[str]:
        sep = "01"
        if s == "":
            return []
        if s == sep:
            return [""]
        out = []
        word = ""
        for j in range(2, len(s), 2):
            i = j - 2
            if s[i:j] != sep:
                word += s[i : i + 1]
            else:
                out.append(word)
                word = ""
        out.append(word)
        return out


s = Solution()
print(s.decode(s.encode(["neet0", "1code", "love", "you"])))

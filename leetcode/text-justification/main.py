from typing import List

class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        line = []
        out = []
        linec = 0

        for cw in range(len(words)):
            if len(line) == 0:
                line.append(words[cw])
                linec += len(words[cw])
                continue

            if linec + len(words[cw]) + 1 <= maxWidth:
                line.append(words[cw])
                linec += len(words[cw]) + 1
                continue

            out.append(self.pad(line, maxWidth, False))
            line = [words[cw]]
            linec = len(words[cw])

        out.append(self.pad(line, maxWidth, True))
        line = [words[cw]]
        linec = len(words[cw])
        
        return out
    
    def pad(self, words: List[str], maxWidth: int, last:bool) -> str:
        # print(words)
        if last or len(words) == 1:
            beginning = " ".join(words)
            total = beginning + " " * (maxWidth - len(beginning))
            return total
        
        emptySpaces = maxWidth - sum(len(s) for s in words)
        
        toFill = len(words) - 1
        d = emptySpaces // toFill
        r = emptySpaces % toFill

        spacings = [(" " * d) for i in range(toFill)]
        for i in range(r):
            spacings[i] = spacings[i] + " "

        out = ""
        for i in range(len(spacings)):
            out += words[i] + spacings[i]

        out += words[toFill]
        return out


s = Solution()
print(s.fullJustify(["This", "is", "an", "example", "of", "text", "justification."], 16))
print(s.fullJustify(["What","must","be","acknowledgment","shall","be"], 16))
print(s.fullJustify(["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], 20))

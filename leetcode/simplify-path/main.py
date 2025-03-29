class Solution:
    
    def simplifyPath(self, path: str) -> str:
        idx = 1
        fullpath = []
        state = "start"
        for (i,c) in enumerate(path):

            if state == "slash" and c != "/":
                idx = i
                state = "inword"
            
            if i == 0:
                continue

            
            elif c != "/" and state == "start":
                idx = i
                state = "inword"

            elif (c == "/" and state == "inword"):
                word = path[idx:i]
                if word == "..":
                    if len(fullpath) != 0:
                        fullpath.pop()
                elif word != ".":
                    fullpath.append(word)
                state = "slash"

            elif (c != "/" and i == len(path) - 1):
                word = path[idx:len(path)]
                if word == "..":
                    if len(fullpath) != 0:
                        fullpath.pop()   
                elif word != ".":
                    fullpath.append(word)
                state = "end"

            elif c != "/" and state == "slash":
                idx = i
                state = "inword"

        return "/" + "/".join(fullpath)
        
s = Solution()
print(s.simplifyPath("/hello/there/"))
print(s.simplifyPath("/hello/there///"))
print(s.simplifyPath("///hello/there/"))
print(s.simplifyPath("/hello///there/"))
print(s.simplifyPath("//hello///there///"))
print(s.simplifyPath("//hello//../there///"))
print(s.simplifyPath("/hello/../there/"))
print(s.simplifyPath("/../../there/"))
print(s.simplifyPath("/home/user/Documents/../Pictures"))
print(s.simplifyPath("/.../a/../b/c/../d/./"))
print(s.simplifyPath("/a/../../b/../c//.//"))
print(s.simplifyPath("/.././GVzvE/./xBjU///../..///././//////T/../../.././zu/q/e"))

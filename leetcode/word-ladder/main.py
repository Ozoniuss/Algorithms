from typing import List
from collections import deque

class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children

class Solution:

    def distance(self, word1, word2):
        c = 0
        for i in range(len(word1)):
            if word1[i] != word2[i]:
                c += 1
        return c
    
    def getNext(self, word):
        n = []
        for i in range(len(word)):
            for j in range(ord('a'), ord('z')+1):
                n.append(word[:i] + chr(j) + word[i+1:])
        return n
    
    def build_masks(self, startWord: str, wordList: List[str]):
        masks = {}
        for word in wordList:
            wordMasks = self.get_masks(word)
            for m in wordMasks:
                if m in masks:
                    masks[m].add(word)
                else:
                    s = set()
                    s.add(word)
                    masks[m] = s
        wordMasks = self.get_masks(startWord)
        for m in wordMasks:
            if m in masks:
                masks[m].add(word)
            else:
                s = set()
                s.add(word)
                masks[m] = s
        return masks
    
    def get_masks(self, word: str):
        masks = []
        for i in range(len(word)):
            mask = word[:i] + '*' + word[i+1:]
            masks.append(mask)
        return masks
        
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        masks = self.build_masks(beginWord, wordList)
        visited = set()
        q = deque()
        q.append((beginWord, 1))

        while len(q) != 0:
            current, d = q[0]
            if current == endWord:
                return d
            visited.add(current)
            q.popleft()
            for m in self.get_masks(current):
                for n in masks[m]:
                    if n not in visited and self.distance(n, current) == 1:
                        q.append((n, d+1))
        
        return 0


    def ladderLengthInneffective(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        remaining = set(wordList)
        q = deque()
        q.append((beginWord, 1))

        while len(q) != 0:
            current, d = q[0]
            if current == endWord:
                return d
            # explored.add(current)
            q.popleft()
            to_remove = []
            for n in remaining:
            # for n in self.getNext(current):
                if self.distance(current, n) == 1:
                    q.append((n, d+1))
                    to_remove.append(n)
            
            for n in to_remove:
                remaining.remove(n)
                

        return 0

s = Solution()
print(s.ladderLength("hit", "cog",["hot","dot","dog","lot","log","cog"]))
print(s.ladderLength("hit", "cog",["hot","dot","dog","lot","log"]))

print(s.ladderLengthInneffective("hit", "cog",["hot","dot","dog","lot","log","cog"]))
print(s.ladderLengthInneffective("hit", "cog",["hot","dot","dog","lot","log"]))

print(len(s.getNext("xxx")))
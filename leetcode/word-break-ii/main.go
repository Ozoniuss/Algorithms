package main

import (
	"fmt"
	"slices"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	// If we got here it means we sliced everything off.
	if len(s) == 0 {
		// empty word so it can for over it
		return []string{""}
	}

	start := 0
	possibleNextIndices := make([]int, 0, 32)
	for next := 1; next <= len(s); next++ {
		if slices.Contains(wordDict, s[start:next]) {
			possibleNextIndices = append(possibleNextIndices, next)
		}
	}
	if len(possibleNextIndices) == 0 {
		return []string{}
	}

	ans := []string{}
	sb := &strings.Builder{}
	for _, ind := range possibleNextIndices {
		currentWord := s[:ind]
		possibilities := wordBreak(s[ind:], wordDict)
		for _, p := range possibilities {
			sb.WriteString(currentWord)
			if len(p) != 0 {
				sb.WriteByte(' ')
				sb.WriteString(p)
			}
			ans = append(ans, sb.String())
			sb.Reset()
		}
	}
	return ans
}

func fmtForEach(in []string) {
	for _, i := range in {
		fmt.Println(i)
	}
	fmt.Println("xxx")
}

func main() {
	fmtForEach(wordBreak("leetcode", []string{"leet", "code", "leetcode"}))
	fmtForEach(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmtForEach(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmtForEach(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

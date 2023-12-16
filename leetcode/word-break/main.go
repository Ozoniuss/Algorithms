package main

import (
	"fmt"
	"slices"
)

func wordBreakCache(s string, wordDict []string, cache *map[string]bool) bool {
	// If we got here it means we sliced everything off.
	if len(s) == 0 {
		return true
	}

	if val, ok := (*cache)[s]; ok {
		return val
	}

	start := 0
	possibleNextIndices := make([]int, 0, 32)
	for next := 1; next <= len(s); next++ {
		if slices.Contains(wordDict, s[start:next]) {
			possibleNextIndices = append(possibleNextIndices, next)
		}
	}
	if len(possibleNextIndices) == 0 {
		return false
	}

	ans := false
	for _, ind := range possibleNextIndices {
		ans = ans || wordBreakCache(s[ind:], wordDict, cache)
	}
	(*cache)[s] = ans
	return ans
}

func wordBreak(s string, wordDict []string) bool {
	cache := make(map[string]bool, 256)
	return wordBreakCache(s, wordDict, &cache)
}

func wordBreakDP(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for next := 1; next <= len(s); next++ {
		for j := next - 1; j >= 0; j-- {
			if dp[j] && slices.Contains(wordDict, s[j:next]) {
				dp[next] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

	fmt.Println(wordBreakDP("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreakDP("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreakDP("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreakDP("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

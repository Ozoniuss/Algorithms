package main

import "fmt"

func numDistinct(s string, t string) int {
	// return count(s, t, 0, 0, make(map[[2]int]int))
	countdpopt(s, t)
	fmt.Println("----")
	countdp(s, t)
	return 0
}

func countdp(s, t string) int {

	if len(t) > len(s) {
		return -1
	}

	// dp(i, j) -> number of seq s[..i] that include t[..j] as a subseq i, j not included
	// essentially we need to compute dp(len(s), len(t))
	//
	// s: ababa
	// t: aa
	/*
		    a a
		  1 0 0
		a 1 1 0
		b 1 1 0
		a 1 2 1
		b 1 2 1
		a 1 3 3

	*/

	// dp(i+1, j) = dp(i,j) if s[i] != t[j]
	// dp(i+1, j+1) = dp(i, j+1) + dp(i, j)
	dp := make([][]int, len(s)+1)
	for i := range len(s) + 1 {
		dp[i] = make([]int, len(t)+1)
	}

	// cannot produce t with an empty string
	for i := range len(t) + 1 {
		dp[0][i] = 0
	}

	// consider that we can produce the null subsequence in exactly one
	// way
	// will also overwrite the above
	for i := range len(s) + 1 {
		dp[i][0] = 1
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := range len(s) + 1 {
		fmt.Println(dp[i])
	}

	return dp[len(s)][len(t)]
}

func countdpopt(s, t string) int {

	if len(t) > len(s) {
		return 0
	}

	// dp(i, j) -> number of seq s[..i] that include t[..j] as a subseq i, j not included
	// essentially we need to compute dp(len(s), len(t))
	//
	// s: ababa
	// t: aa
	/*
		    a a
		  1 0 0
		a 1 1 0
		b 1 1 0
		a 1 2 1
		b 1 2 1
		a 1 3 3

	*/

	// dp(i+1, j) = dp(i,j) if s[i] != t[j]
	// dp(i+1, j+1) = dp(i, j+1) + dp(i, j)

	// we only need the last row as an optimized version
	dp := make([]int, len(t)+1)

	// cannot produce t with an empty string
	for i := range len(t) + 1 {
		dp[i] = 0
	}

	// starting point for empty source target string
	dp[0] = 1

	for i := 1; i < len(s)+1; i++ {
		fmt.Println(dp)
		prev := dp[0]
		for j := 1; j < min(i+1, len(t)+1); j++ {
			p := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] += prev
			}
			prev = p
		}
	}
	// fmt.Println(dp)

	return dp[len(t)]
}

func count(s, t string, is int, it int, cache map[[2]int]int) int {

	if v, ok := cache[[2]int{is, it}]; ok {
		return v
	}

	// is -> current char in s
	// it -> next char that I need to fill in t

	// I have filled all of t
	if it == len(t) {
		return 1
	}
	if is > len(s) {
		return 0
	}

	// start from i
	tot := 0
	for i := is; i < len(s); i++ {
		if is == 0 && it == 0 {
			fmt.Println(i)
		}

		// choose i as the next character of t
		if s[i] == t[it] {
			tot += count(s, t, i+1, it+1, cache)
		}
	}
	cache[[2]int{is, it}] = tot
	return tot
}

func main() {
	// fmt.Println(numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("ababa", "aa"))
}

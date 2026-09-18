package main

import "fmt"

// aabbcc aabbcc a*c a?*c aab* aab??? a*c* *c
// a*

// ismatch()
// ismatch(i, j) = ismatch(i-1, j-1) if s[i] = p[j] or p[j] == ?
// ismatch(i, j) = ismatch(i-1) if p[j] == *
//      a * c
//    1 0 0 0
//  a 0 1 1 0
//  a 0 0 1 0
//  b 0 0 1 0
//  b 0 0 1 0
//  c 0 0 1 1
//  c 0 0 1 1

//      a ? c
//    1 0 0 0
//  a 0 1 0 0
//  b 0 0 1 0
//  b 0 0 0 0

//      a a * c
//    1 0 0 0 0
//  a 0 1 0 0 0
//  a 0 0 1 1 0
//  b 0 0 0 1 0
//  b 0 0 0 1 0
//  c 0 0 0 1 1
//  c 0 0 0 1 1

//      a b * c
//    1 0 0 0 0
//  a 0 1 0 0 0
//  a 0 0 0 0 0
//  b 0 0 0 0 0
//  b 0 0 0 0 0
//  c 0 0 0 0 0
//  c 0 0 0 0 0

//      *
//    1 0
//  a 0 1
//  a 0 0
//  b 0 0
//  b 0 0
//  c 0 0
//  c 0 0

//      * a * c
//    1 0 0 0 0
//  a 0 1 1 1 0
//  d 0 0 1 1 0
//  c 0 0 0 1 0
//  e 0 0 0 1 0
//  b 0 0 0 1 1
//  c 0 0 0 1 1

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range len(s) + 1 {
		dp[i] = make([]bool, len(p)+1)
	}
	// nothing matches nothing
	dp[0][0] = true

	// wildcard in the beginning can match empty words
	offset := 0
	for offset < len(p) && p[offset] == '*' {
		dp[0][offset+1] = true
		offset++
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			switch p[j-1] {
			case '*':
				// Either match either a new char or no char with the
				// star for the first time (thus using the star and
				// advancing j) or reuse the starß
				dp[i][j] = dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]
			case '?':
				dp[i][j] = dp[i-1][j-1]
			default:
				if s[i-1] == p[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	for _, line := range dp {
		fmt.Println(line)
	}
	return dp[len(s)][len(p)]
}

func isMatch2(s string, p string) bool {
	dp := make([]bool, len(p)+1)
	// nothing matches nothing
	dp[0] = true

	// wildcard in the beginning can match empty words
	offset := 0
	for offset < len(p) && p[offset] == '*' {
		dp[offset+1] = true
		offset++
	}

	for i := 1; i < len(s)+1; i++ {
		prev := dp[0] // prev is basically dp[i-1][j-1]
		for j := 1; j < len(p)+1; j++ {
			aux := dp[j]
			dp[j] = false
			switch p[j-1] {
			case '*':
				// Either match either a new char or no char with the
				// star for the first time (thus using the star and
				// advancing j) or reuse the starß
				dp[j] = prev || dp[j-1] || aux
			case '?':
				dp[j] = prev
			default:
				if s[i-1] == p[j-1] {
					dp[j] = prev
				}
			}
			prev = aux
		}
		// in future iterations this is actually 0 and represents
		// pattern with empty word
		dp[0] = false
	}

	// for _, line := range dp {
	// 	fmt.Println("dp", line)
	// }
	return dp[len(p)]
}

func main() {
	fmt.Println(isMatch("aabbcc", "aabbcc"))
	fmt.Println(isMatch("aabbcc", "aabbcd"))
	fmt.Println(isMatch("aabbcc", "aabb"))
	fmt.Println(isMatch("aabbcc", "aabb?c"))
	fmt.Println(isMatch("aabbcc", "aabbc?"))
	fmt.Println(isMatch("aabbcc", "aabb??"))
	fmt.Println(isMatch("aabbcc", "?abb??"))
	fmt.Println(isMatch("aabbcc", "?abb??"))
	fmt.Println(isMatch("aabbcc", "a*"))
	fmt.Println(isMatch("aabbcc", "*c"))
	fmt.Println(isMatch("aabbcc", "a*c"))
	fmt.Println(isMatch("aabbcc", "a**c"))
	fmt.Println(isMatch("aabbcc", "a**b"))
	fmt.Println(isMatch("aabbcc", "*******"))
	fmt.Println(isMatch("aabbcc", "*cb*"))
	fmt.Println(isMatch("aabbcc", "*bc*"))
	fmt.Println(isMatch("aabbcc", "*??*"))
	fmt.Println(isMatch("aabbcc", "*???????*"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("adceb", "**a*b"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println("===")
	fmt.Println(isMatch2("aabbcc", "aabbcc"))
	fmt.Println(isMatch2("aabbcc", "aabbcd"))
	fmt.Println(isMatch2("aabbcc", "aabb"))
	fmt.Println(isMatch2("aabbcc", "aabb?c"))
	fmt.Println(isMatch2("aabbcc", "aabbc?"))
	fmt.Println(isMatch2("aabbcc", "aabb??"))
	fmt.Println(isMatch2("aabbcc", "?abb??"))
	fmt.Println(isMatch2("aabbcc", "?abb??"))
	fmt.Println(isMatch2("aabbcc", "a*"))
	fmt.Println(isMatch2("aabbcc", "*c"))
	fmt.Println(isMatch2("aabbcc", "a*c"))
	fmt.Println(isMatch2("aabbcc", "a**c"))
	fmt.Println(isMatch2("aabbcc", "a**b"))
	fmt.Println(isMatch2("aabbcc", "*******"))
	fmt.Println(isMatch2("aabbcc", "*cb*"))
	fmt.Println(isMatch2("aabbcc", "*bc*"))
	fmt.Println(isMatch2("aabbcc", "*??*"))
	fmt.Println(isMatch2("aabbcc", "*???????*"))
	fmt.Println(isMatch2("adceb", "*a*b"))
	fmt.Println(isMatch2("adceb", "**a*b"))
	fmt.Println(isMatch2("aa", "a"))
}

package main

import (
	"fmt"
)

// aabbcc aabbcc a*c a?*c aab* aab??? a*c* *c
// a*

//      a * c
//    1 0 0 0
//  a 0 1 1 0
//  a 0 0 1 0
//  b 0 0 1 0
//  b 0 0 1 0
//  c 0 0 1 1
//  c 0 0 1 1

//      a . * c
//    1 0 0 0 0
//  a 0 1 0 0 0
//  a 0 0 1 1 0
//  b 0 0 0 1 0
//  c 0 0 0 1 1

//      . * . *
//    1 0 0 0 0
//  a 0 1 1 0 0
//  a 0 0 1 1 1
//  b 0 0 1 1 1
//  c 0 0 1 1 1

//     a*
//   1 0
// a 0 1
// a 0 1

//     a .
//   1 0 0
// a 0 1 0
// a 0 0 1

// s = aab, p = aa*b

//      a  a* b
//    1 0  0  0
//  a 0 1  1  0
//  a 0 0  1  0
//  a 0 0  1  0
//  b 0 0  0  1

//      c* a* b
//    1 1  1  0
//  a 0 0  1  0
//  a 0 0  1  0
//  b 0 0  0  1

func isMatch(s string, p string) bool {

	adjp := []byte{}
	wildcard := []bool{}

	skip := false
	for i := range len(p) {
		if i == 0 || skip {
			skip = false
			continue
		}
		if p[i] != '*' {
			adjp = append(adjp, p[i-1])
			wildcard = append(wildcard, false)
		} else {
			adjp = append(adjp, p[i-1])
			wildcard = append(wildcard, true)
			i = i + 1
			skip = true
		}
	}
	if p[len(p)-1] != '*' {
		adjp = append(adjp, p[len(p)-1])
		wildcard = append(wildcard, false)
	}

	// fmt.Println("adjp", string(adjp))
	// fmt.Println("wildcard", wildcard)
	p = string(adjp)

	dp := make([][]bool, len(s)+1)
	for i := range len(s) + 1 {
		dp[i] = make([]bool, len(p)+1)
	}
	// nothing matches nothing
	dp[0][0] = true

	// match empty chars with first wildcards
	isWildcard := true
	for i := range len(wildcard) {
		isWildcard = wildcard[i]
		if !isWildcard {
			break
		}
		dp[0][i+1] = true
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if wildcard[j-1] {
				useWildcard := dp[i-1][j-1]
				if p[j-1] != '.' {
					useWildcard = dp[i-1][j-1] && s[i-1] == p[j-1]
				}
				advanceWildcard := dp[i-1][j]
				if p[j-1] != '.' {
					advanceWildcard = dp[i-1][j] && s[i-1] == p[j-1]
				}
				// will always be able to match no character
				dp[i][j] = useWildcard || advanceWildcard || dp[i][j-1]
			} else {
				if p[j-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				} else if s[i-1] == p[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	// for _, line := range dp {
	// 	fmt.Println(line)
	// }
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
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("aabc", ".*.*"))
	fmt.Println(isMatch("aab", "c*a*b"))
}

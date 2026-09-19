package main

import "fmt"

/*
    r o s
  0 1 2 3
h 1 1 2 3
o 2 2 1 2
r 3 2 2 2
s 4 3 3 2
e 5 4 4 3


    r o s
  0 1 2 3
h 1 1 2 3
r 2 1 2 3
o 3 2 1 2

*/

// ...a ->  ...b (N)
//
// ...ax -> ...by (N+1)
// ...ax -> ...bx (N)
// ...ax -> ...b (N+1)
// ...a  -> ...bx (N+1)

func minDistance(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)
	for i := range len(word1) + 1 {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := range len(word1) + 1 {
		dp[i][0] = i
	}
	for j := range len(word2) + 1 {
		dp[0][j] = j
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// honestly I'm not sure I understand why this works.
				// but it works. feels right.
				//
				//  ...xA ....yB (x and y may be the same char)
				//  previous problem: ...x -> ...y (I turned A into B)
				//  previous problem: ...x ->  ...yB (I removed A)
				//  previous problem: ...xA -> ...y (I added B)
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i][j-1]+1, dp[i-1][j]+1)
			}
		}
	}
	for _, line := range dp {
		fmt.Println(line)
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

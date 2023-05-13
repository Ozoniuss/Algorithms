package main

import (
	"fmt"
)

// Robot needs to reach from (0,0) to (m,n)
// Note that the number of paths to (m,n) is equal to number of paths to
// (m-1, n) + number of paths to (m, n-1)

// Note that this can be solved easily using a mathematical formula, by just
// computing the number of ways to choose m-1 down moves and n-1 right moves
// out of m-n-2 required moves. So it's just comb(m-n-2, n-1)

// m means lines, n means columns
func uniquePaths(m int, n int) int {
	paths := make([][]int, m, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n, n)
	}
	paths[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				paths[i][j] += paths[i-1][j]
			}
			if j-1 >= 0 {
				paths[i][j] += paths[i][j-1]
			}
		}
	}

	return paths[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}

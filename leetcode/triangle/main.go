package main

func minimumTotal(triangle [][]int) int {
	N := len(triangle)
	dp := make([][]int, N)
	for i := range N {
		dp[i] = make([]int, len(triangle[i]))
	}

	for j := range len(triangle[N-1]) {
		dp[N-1][j] = triangle[N-1][j]
	}

	for i := N - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func main() {

}

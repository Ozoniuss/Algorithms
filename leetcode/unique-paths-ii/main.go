package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	paths := make([][]int, m, n)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				paths[i][j] = -1
			}
		}
	}
	paths[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if paths[i][j] == -1 {
				continue
			}
			if i-1 >= 0 && paths[i-1][j] != -1 {
				paths[i][j] += paths[i-1][j]
			}
			if j-1 >= 0 && paths[i][j-1] != -1 {
				paths[i][j] += paths[i][j-1]
			}
		}
	}

	return paths[m-1][n-1]
}

func main() {

}

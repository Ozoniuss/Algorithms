package main

func countServers(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	// count every row that has at least two servers to exclude them
	// from count next time
	seen := make([]bool, R)
	total := 0
	for i := range R {
		cnt := 0
		for j := range C {
			if grid[i][j] == 0 {
				continue
			}
			cnt += 1
		}
		if cnt > 1 {
			seen[i] = true
			total += cnt
		}
	}

	for j := range C {
		cnt := 0
		adj := 0
		for i := range R {
			if grid[i][j] == 0 {
				continue
			}
			cnt += 1
			if seen[i] {
				adj -= 1
			}
		}
		if cnt > 1 {
			total += cnt + adj
		}
	}
	return total
}

func main() {

}

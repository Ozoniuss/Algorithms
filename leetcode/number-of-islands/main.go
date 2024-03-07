package main

// I'm quite proud of this. I solved it while being moderately drunk and
// talking to my gf at the same time.

import "fmt"

func numIslands(grid [][]byte) int {
	discovered := make(map[[2]int]struct{})
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// explore this space to find the island
			total += explore(i, j, discovered, grid)
		}
	}
	return total
}

func explore(x, y int, discovered map[[2]int]struct{}, grid [][]byte) int {

	// if this was not discovered before it's going to be used to generate a new
	// map.This is then obviously returning 1.
	if _, ok := discovered[[2]int{x, y}]; ok {
		return 0
	}

	fmt.Println("lol", x, y)

	// no land
	if grid[x][y] == '0' {
		return 0
	}

	exploreAndMark(x, y, discovered, grid)
	return 1
}

func exploreAndMark(x, y int, discovered map[[2]int]struct{}, grid [][]byte) {

	if x < 0 || x >= len(grid) {
		return
	}
	if y < 0 || y >= len(grid[0]) {
		return
	}

	if grid[x][y] == '0' {
		discovered[[2]int{x, y}] = struct{}{}
		return
	}

	if _, ok := discovered[[2]int{x, y}]; ok {
		return
	}
	discovered[[2]int{x, y}] = struct{}{}

	nextPos := [][2]int{{x, y + 1}, {x, y - 1}, {x - 1, y}, {x + 1, y}}

	for _, next := range nextPos {
		exploreAndMark(next[0], next[1], discovered, grid)
	}
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}

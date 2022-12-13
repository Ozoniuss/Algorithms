package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLine reads a line from the input file and appends it as a row with
// integers to a matrix.
func readLine(line string, matrix *[][]int) {
	row := []int{}
	for _, c := range line {
		if c == 'S' {
			row = append(row, 'a')
		} else if c == 'E' {
			row = append(row, 'z')
		} else {
			row = append(row, int(c))
		}
	}
	(*matrix) = append((*matrix), row)
}

// getNeightbours returns all neighbours of a point.
func getNeighbours(point [2]int, width, height int) [][2]int {
	neighbours := make([][2]int, 0, 4)
	if point[0] > 0 {
		neighbours = append(neighbours, [2]int{point[0] - 1, point[1]})
	}
	if point[0] < width-1 {
		neighbours = append(neighbours, [2]int{point[0] + 1, point[1]})
	}
	if point[1] > 0 {
		neighbours = append(neighbours, [2]int{point[0], point[1] - 1})
	}
	if point[1] < height-1 {
		neighbours = append(neighbours, [2]int{point[0], point[1] + 1})
	}
	return neighbours
}

func canClimb(prev, next int) bool {
	return next-prev <= 1
}

// findShortestPath is a backtracking approach that generates all solutions
// and returns the shortest one.
//
// Note that it doesn't make sense to walk through a location twice. If a path
// has a loop, there exists a shorter path without that loop.
func findShortestPath(matrix [][]int, src [2]int, dst [2]int) {
	visited := make(map[[2]int]struct{})
	//q := make([][2]int, 0, cap(matrix))

	visited[src] = struct{}{}
	q := make([][2]int, 0, len(matrix))

	// Add the starting point to the bfs queue.
	q = append(q, src)

	bfs(matrix, dst, &visited, &q)
}

// findShortestPathMultiSource is the same as findShortestPath, except that it
// accepts multiple sources as input, and finds the shortest path starting
// from any of these sources.
func findShortestPathMultiSource(matrix [][]int, src [][2]int, dst [2]int) {
	visited := make(map[[2]int]struct{})
	//q := make([][2]int, 0, cap(matrix))

	q := make([][2]int, 0, len(matrix))

	// Add the starting points to the bfs queue.
	for _, s := range src {
		q = append(q, s)
	}

	bfs(matrix, dst, &visited, &q)
}

// bfs performs bread-first search traversal on the map.
func bfs(matrix [][]int, dst [2]int, visited *map[[2]int]struct{}, q *[][2]int) {

	currentDepth := 0

	for len(*q) > 0 {

		fmt.Printf("Current depth is %d\n", currentDepth)

		nextQ := make([][2]int, 0, len(matrix))
		for len(*q) > 0 {

			currentVertex := (*q)[0]

			if currentVertex == dst {
				fmt.Printf("Reached current verted at depth %d\n", currentDepth)
				return
			}

			(*q) = (*q)[1:]

			for _, n := range getNeighbours(currentVertex, len(matrix), len(matrix[0])) {
				if _, ok := (*visited)[n]; ok == false {
					if canClimb(matrix[currentVertex[0]][currentVertex[1]], matrix[n[0]][n[1]]) {
						nextQ = append(nextQ, n)
						(*visited)[n] = struct{}{}
					}
				}
			}
		}

		*q = nextQ
		currentDepth += 1

	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, 0, 100)
	src := [2]int{}
	dst := [2]int{}

	srcs := make([][2]int, 0, len(matrix))

	part := 2

	scanner := bufio.NewScanner(f)
	lineNumber := 0

	if part == 1 {
		for scanner.Scan() {
			line := scanner.Text()

			if pos := strings.Index(line, "S"); pos != -1 {
				src = [2]int{lineNumber, pos}
			}

			if pos := strings.Index(line, "E"); pos != -1 {
				dst = [2]int{lineNumber, pos}
			}

			readLine(line, &matrix)
			lineNumber += 1
		}

		findShortestPath(matrix, src, dst)
	} else if part == 2 {
		for scanner.Scan() {
			line := scanner.Text()

			// Determines all starting positions, which can either be S or a.
			for pos, c := range line {
				if c == 'a' || c == 'S' {
					srcs = append(srcs, [2]int{lineNumber, pos})
				}
			}

			if pos := strings.Index(line, "E"); pos != -1 {
				dst = [2]int{lineNumber, pos}
			}

			readLine(line, &matrix)
			lineNumber += 1
		}

		findShortestPathMultiSource(matrix, srcs, dst)
	}
}

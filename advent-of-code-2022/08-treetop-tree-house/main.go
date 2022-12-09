package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction byte

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

// readTrees reads the trees from a line and adds them as an array to the
// provided tree matrix.
func readTrees(line string, treeMatrix *[][]int) error {

	trees := make([]int, 0, len(line))
	for _, l := range line {
		val, err := strconv.ParseInt(string(l), 10, 32)
		if err != nil {
			return err
		}
		trees = append(trees, int(val))
	}
	*treeMatrix = append(*treeMatrix, trees)
	return nil
}

// findVisible finds all trees that are visible from some direction and adds
// their position to the set with visible elements.
func findVisible(treeMatrix *[][]int, visible *map[[2]int]int) {

	// read like: -->
	for i := 0; i < len(*treeMatrix); i++ {
		tallestTree := -1
		for j := 0; j < len((*treeMatrix)[0]); j++ {
			if (*treeMatrix)[i][j] > tallestTree {
				(*visible)[[2]int{i, j}] = (*treeMatrix)[i][j]
				tallestTree = (*treeMatrix)[i][j]
			}
		}
	}

	// read like: <--
	for i := 0; i < len(*treeMatrix); i++ {
		tallestTree := -1
		for j := len((*treeMatrix)[0]) - 1; j >= 0; j-- {
			if (*treeMatrix)[i][j] > tallestTree {
				(*visible)[[2]int{i, j}] = (*treeMatrix)[i][j]
				tallestTree = (*treeMatrix)[i][j]
			}
		}
	}

	// read like: |
	//            v
	for i := 0; i < len(*treeMatrix); i++ {
		tallestTree := -1
		for j := 0; j < len((*treeMatrix)[0]); j++ {
			if (*treeMatrix)[j][i] > tallestTree {
				(*visible)[[2]int{j, i}] = (*treeMatrix)[j][i]
				tallestTree = (*treeMatrix)[j][i]
			}
		}
	}

	// read like: ^
	//            |
	for i := 0; i < len(*treeMatrix); i++ {
		tallestTree := -1
		for j := len((*treeMatrix)[0]) - 1; j >= 0; j-- {
			if (*treeMatrix)[j][i] > tallestTree {
				(*visible)[[2]int{j, i}] = (*treeMatrix)[j][i]
				tallestTree = (*treeMatrix)[j][i]
			}
		}
	}

}

// findScenicScore finds the scenic score of tree at position i and j.
func findScenicScore(pos_i int, pos_j int, treeMatrix *[][]int) int {
	left := 0
	right := 0
	up := 0
	down := 0

	height := (*treeMatrix)[pos_i][pos_j]

	for i := pos_i - 1; i >= 0; i-- {
		up += 1
		if (*treeMatrix)[i][pos_j] >= height {
			break
		}
	}

	for i := pos_i + 1; i < len(*treeMatrix); i++ {
		down += 1
		if (*treeMatrix)[i][pos_j] >= height {
			break
		}
	}

	for j := pos_j - 1; j >= 0; j-- {
		left += 1
		if (*treeMatrix)[pos_i][j] >= height {
			break
		}
	}

	for j := pos_j + 1; j < len((*treeMatrix)[0]); j++ {
		right += 1
		if (*treeMatrix)[pos_i][j] >= height {
			break
		}
	}

	return left * right * up * down
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	treeMatrix := make([][]int, 0)
	visible := make(map[[2]int]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		readTrees(line, &treeMatrix)
	}
	findVisible(&treeMatrix, &visible)
	fmt.Printf("There are a total of %d visible trees.\n", len(visible))

	maxScenic := 0
	for i := 0; i < len(treeMatrix); i++ {
		for j := 0; j < len(treeMatrix[0]); j++ {
			if val := findScenicScore(i, j, &treeMatrix); val > maxScenic {
				maxScenic = val
			}
		}
	}

	fmt.Printf("The maximum scenic score of a tree is %d\n", maxScenic)
}

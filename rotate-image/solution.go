package main

import "fmt"

// 90 degree rotation, (i,j) -> (j, n-1-i)

// rotateSingle moves the current value to its rotated location, and
// returns the old value from that location and the positions
func rotateSingle(matrix *[][]int, current_val, r, c, size int) (int, int, int) {
	value := (*matrix)[c][size-1-r]
	(*matrix)[c][size-1-r] = current_val
	return c, size - 1 - r, value
}

// fullRotation performs 4 rotations of the matrix, which allows to
// rotate 4 of the values on the same outer square with constant memory
func fullRotation(matrix *[][]int, r, c, size int) {
	current_val := (*matrix)[r][c]
	for idx := 0; idx < 4; idx++ {
		r, c, current_val = rotateSingle(matrix, current_val, r, c, size)
	}
}

// Complexity:  ~ Th(n^2)
// Extra memory: Th(1)
func rotate(matrix [][]int) {
	size := len(matrix)
	if size == 1 {
		return
	}
	for r := 0; r < int(size/2); r++ {
		for c := r; c < size-r-1; c++ {
			fullRotation(&matrix, r, c, size)
		}
	}
}

func main() {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix)
}

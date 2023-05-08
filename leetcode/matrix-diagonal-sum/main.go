package main

func diagonalSum(mat [][]int) int {
	sum := 0
	size := len(mat)

	// Helps not have this if inside the for loop.
	if size%2 == 1 {
		sum += mat[size/2][size/2]
	}

	for i := 0; i < size/2; i++ {
		sum += (mat[i][i] + mat[size-i-1][size-i-1])
		sum += (mat[i][size-i-1] + mat[size-i-1][i])
	}

	return sum
}

func main() {

}

package main

import "fmt"

func generateMatrix(n int) [][]int {

	L := n
	C := n
	boundaries := [4]int{0, L, 0, C}

	out := make([][]int, n, n)
	for i := range n {
		out[i] = make([]int, n, n)
	}
	c := 1

	for {
		for j := boundaries[2]; j < boundaries[3]; j++ {
			i := boundaries[0]
			out[i][j] = c
			c += 1
		}
		boundaries[0] += 1
		if boundaries[0] == boundaries[1] {
			return out
		}
		for i := boundaries[0]; i < boundaries[1]; i++ {
			j := boundaries[3] - 1
			out[i][j] = c
			c += 1
		}
		boundaries[3] -= 1
		if boundaries[2] == boundaries[3] {
			return out
		}
		for j := boundaries[3] - 1; j >= boundaries[2]; j-- {
			i := boundaries[1] - 1
			out[i][j] = c
			c += 1
		}
		boundaries[1] -= 1
		if boundaries[0] == boundaries[1] {
			return out
		}
		for i := boundaries[1] - 1; i >= boundaries[0]; i-- {
			j := boundaries[2]
			out[i][j] = c
			c += 1
		}
		boundaries[2] += 1
		if boundaries[2] == boundaries[3] {
			return out
		}
	}
}

func main() {
	fmt.Println(generateMatrix(3))
}

package main

import "fmt"

/*

010
001
100

0,1
1,2
2,0

1,2
2,2


000
001
001

*/

// I didn't come up with this, but this is such a clever idea. Not sure
// I would have ever come up with something like this.

func largestOverlap(img1 [][]int, img2 [][]int) int {
	i1 := [][2]int{}
	i2 := [][2]int{}

	cnt := map[[2]int]int{}

	for i := range len(img1) {
		for j := range len(img1[0]) {
			if img1[i][j] == 1 {
				i1 = append(i1, [2]int{i, j})
			}
		}
	}

	for i := range len(img2) {
		for j := range len(img2[0]) {
			if img2[i][j] == 1 {
				i2 = append(i2, [2]int{i, j})
			}
		}
	}

	for _, x := range i1 {
		for _, y := range i2 {
			cnt[[2]int{y[0] - x[0], y[1] - x[1]}] += 1
		}
	}

	m := 0
	for _, v := range cnt {
		if v > m {
			m = v
		}
	}

	return m
}

func main() {
	n1 := 0b011011111101
	n2 := 0b111101001010
	fmt.Println(n1, n2)
}

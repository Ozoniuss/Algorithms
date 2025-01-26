package main

import "fmt"

type loc = [2]int

var (
	UP    = loc{-1, 0}
	DOWN  = loc{1, 0}
	LEFT  = loc{0, -1}
	RIGHT = loc{0, +1}
)

func nextdir(cdir loc) loc {
	if cdir == UP {
		return RIGHT
	}
	if cdir == RIGHT {
		return DOWN
	}
	if cdir == DOWN {
		return LEFT
	}
	if cdir == LEFT {
		return UP
	}
	panic("not gonna happen")
}

func spiralOrder(matrix [][]int) []int {

	L := len(matrix)
	C := len(matrix[0])

	dir := RIGHT
	cloc := loc{0, 0}
	out := []int{}

	visited := make(map[loc]struct{}, L*C)
	visited[cloc] = struct{}{}
	out = append(out, matrix[cloc[0]][cloc[1]])

	for len(visited) < L*C {
		n := loc{cloc[0] + dir[0], cloc[1] + dir[1]}
		_, ok := visited[n]
		if ok || (n[0] < 0 || n[0] >= L || n[1] < 0 || n[1] >= C) {
			dir = nextdir(dir)
			continue
		} else {
			visited[n] = struct{}{}
			out = append(out, matrix[n[0]][n[1]])
			cloc = n
		}
	}
	return out
}

func main() {
	m := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}
	fmt.Println(spiralOrder(m))
}

package main

func exist(board [][]byte, word string) bool {

	R, C := len(board), len(board[0])
	for i := range R {
		for j := range C {
			f := explore(board, word, i, j, 0, make(map[[2]int]struct{}))
			if f {
				return true
			}
		}
	}

	return false
}

func explore(board [][]byte, word string, x int, y int, idx int, visited map[[2]int]struct{}) bool {

	R, C := len(board), len(board[0])
	// cannot visit more than once
	if _, ok := visited[[2]int{x, y}]; ok {
		return false
	}

	if board[x][y] != word[idx] {
		return false
	}

	visited[[2]int{x, y}] = struct{}{}
	if idx+1 == len(word) {
		return true
	}
	next := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, n := range next {
		i, j := n[0], n[1]
		if !(0 <= i && i < R && 0 <= j && j < C) {
			continue
		}
		found := explore(board, word, i, j, idx+1, visited)
		if found {
			return true
		}
	}

	delete(visited, [2]int{x, y})
	return false
}

func main() {

}

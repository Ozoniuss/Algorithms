package main

import (
	"fmt"
	"slices"
)

type stateEl struct {
	left      int
	pos       [2]int
	available [10]int
}

func solveSudoku(board [][]byte) {

	R, C := len(board), len(board[0])
	states := []stateEl{}
	for i := range R {
		for j := range C {
			if board[i][j] != '.' {
				continue
			}
			fmt.Println("i,j", i, j)
			available := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
			for ii := range 9 {
				c := board[ii][j]
				if c != '.' {
					available[c-'0'] = 0
				}
			}
			for jj := range 9 {
				c := board[i][jj]
				if c != '.' {
					available[c-'0'] = 0
				}
			}

			sqi := (i / 3) * 3
			sqj := (j / 3) * 3
			for x := range 3 {
				for y := range 3 {
					ii := sqi + x
					jj := sqj + y
					if c := board[ii][jj]; c != '.' {
						available[c-'0'] = 0
					}
				}
			}
			left := 9
			for x := 1; x < 10; x++ {
				if available[x] == 0 {
					left--
				}
			}
			el := stateEl{
				available: available,
				pos:       [2]int{i, j},
				left:      left,
			}
			states = append(states, el)
		}
	}
	slices.SortStableFunc(states, func(a stateEl, b stateEl) int {
		return a.left - b.left
	})
	b := explore(board, states)
	board = b
}

func explore(board [][]byte, states []stateEl) [][]byte {
	fmt.Println("len", len(states), states)
	if len(states) == 0 {
		return board
	}
	top := states[0]
	if top.left == 0 {
		return nil
	}
	statesc := slices.Clone(states[1:])

	// this is the cell that we want to fill
	i, j := top.pos[0], top.pos[1]

	if board[i][j] != '.' {
		panic("sth went wrong")
	}
	// try out all available places
	for x := 1; x < 10; x++ {
		// will write x if it was ever available, and then it will no
		// longer be available for the columns, rows or squares that
		// interact with i, j
		if top.available[x] == 1 {
			board[i][j] = byte(x) + '0'
			// recalculate states
			statescl := slices.Clone(statesc)
			for sidx := range len(statescl) {
				// wasn't available in the first place for this other
				// cell
				if statescl[sidx].available[x] == 0 {
					continue
				}
				// row
				if statescl[sidx].pos[0] == i {
					statescl[sidx].available[x] = 0
					statescl[sidx].left--

				} else if statescl[sidx].pos[1] == j {
					statescl[sidx].available[x] = 0
					statescl[sidx].left--

				} else {
					// square
					sqi := (i / 3) * 3
					sqj := (j / 3) * 3
					if statescl[sidx].pos[0]-sqi <= 2 && statescl[sidx].pos[1]-sqj <= 2 && statescl[sidx].pos[0]-sqi >= 0 && statescl[sidx].pos[1]-sqj >= 0 {
						statescl[sidx].available[x] = 0
						statescl[sidx].left--

					}
				}
			}
			slices.SortStableFunc(statescl, func(a stateEl, b stateEl) int {
				return a.left - b.left
			})
			b := explore(board, statescl)
			if b != nil {
				return b
			}

			board[i][j] = '.'
		}
	}
	return nil

}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}

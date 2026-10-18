package main

import (
	"fmt"
	"slices"
	"strings"
)

func findWords(board [][]byte, words []string) []string {
	slices.Sort(words)

	// note that words are unique. so if a word is in words, its value in prefix
	// targets must be [p, p+1]
	prefixTargets := allPrefixTargets(words)
	wordMap := make(map[string]struct{})
	for _, w := range words {
		wordMap[w] = struct{}{}
	}
	found := make(map[string]struct{})

	R, C := len(board), len(board[0])
	for i := range R {
		for j := range C {
			// explore every cell and find all words that can be reached from there
			explore(board, R, C, words, wordMap, make([]byte, 0, 16), i, j, prefixTargets, found)
		}
	}

	out := make([]string, 0, len(found))
	for k := range found {
		out = append(out, k)
	}

	return out
}

// explore starts from a board state and keeps exploring
func explore(board [][]byte, R int, C int, words []string, wordMap map[string]struct{}, cword []byte, x int, y int, prefixTarget map[string][2]int, found map[string]struct{}) {

	istate := board[x][y]
	// set state
	board[x][y] = '*'

	defer func() {
		board[x][y] = istate
	}()

	// we can actually pass this as a copy to future callers and it will still
	// reuse the backing slice, but since cword copies the slice's header then
	// basically each neighbour will actually end up overwriting the string
	// at the correct location. lol.
	cword = append(cword, istate)

	// find how many possible words are there to explore.
	pos, ok := prefixTarget[string(cword)]
	// no targets at all
	if !ok {
		return
	}

	// we have found a word
	if _, ok := wordMap[string(cword)]; ok {
		found[string(cword)] = struct{}{}

		// in this situation, this is the only word, so we can exit early.
		if pos[1]-pos[0] == 1 {
			return
		}
	}

	// keep exploring
	next := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	for _, n := range next {
		i, j := n[0], n[1]
		// skip invalid neighbours
		if !(0 <= i && i < R && 0 <= j && j < C) {
			continue
		}
		// skip neighbours that have already been visited
		if board[i][j] == '*' {
			continue
		}
		explore(board, R, C, words, wordMap, cword, i, j, prefixTarget, found)
	}

}

// words must be sorted
func prefixTarget(words []string, prefix string) (int, int) {
	start, _ := slices.BinarySearchFunc(words, prefix, strings.Compare)

	end, _ := slices.BinarySearchFunc(words, prefix, func(word, prefix string) int {
		if strings.HasPrefix(word, prefix) {
			return -1
		}
		return strings.Compare(word, prefix)
	})

	return start, end
}

// words must be sorted
func allPrefixTargets(words []string) map[string][2]int {
	prefixTargets := make(map[string][2]int)

	for widx, w := range words {
		for idx := range len(w) + 1 {
			subw := w[:idx]
			if pt, ok := prefixTargets[subw]; !ok {
				prefixTargets[subw] = [2]int{widx, widx + 1}
			} else {
				prefixTargets[subw] = [2]int{pt[0], pt[1] + 1}
			}
		}
	}

	return prefixTargets
}

func main() {
	words := []string{"abcde", "abcdf", "abdf", "abxy", "ad", "adf", "bf"}

	prefixTargets := make(map[string][2]int)

	for widx, w := range words {
		for idx := range len(w) + 1 {
			subw := w[:idx]
			if pt, ok := prefixTargets[subw]; !ok {
				prefixTargets[subw] = [2]int{widx, widx + 1}
			} else {
				prefixTargets[subw] = [2]int{pt[0], pt[1] + 1}
			}
		}
	}

	fmt.Println(prefixTargets)

	fmt.Println(prefixTarget(words, "abd"))
}

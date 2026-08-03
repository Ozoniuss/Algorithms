package main

import "fmt"

func stoneGameIII(piles []int) string {
	cache := make(map[int]int, 0)
	m := maxDiffForFirstPlayer(piles, 0, cache)
	if m == 0 {
		return "Tie"
	}
	if m > 0 {
		return "Alice"
	}
	return "Bob"
}

// I actually ended up looking up how they do it since I got so lost in my recursion,
// it's much easier to see this as a diff assuming the first player starts rather than
// a score, since a score requires knowing a partial score for both players and it can
// be difficult to compute that
//
// so let's just calculate the maximum diff (first player score - last player score) that
// we can obtain from lidx assuming we start as the first player with no diff
func maxDiffForFirstPlayer(piles []int, lidx int, cache map[int]int) int {

	if len(piles) < lidx {
		panic("wtf")
	}

	// cache asssumes currentStones was 0 when computing
	if c, ok := cache[lidx]; ok {
		return c
	}

	if lidx == len(piles) {
		return 0
	}

	// 1 stone left
	if lidx == len(piles)-1 {
		return piles[lidx]
	}

	// 2 stones left
	if lidx == len(piles)-2 {
		// in the first case, the second element goes to the second player so
		// adjust the diff
		opt1 := piles[lidx] - piles[lidx+1]
		opt2 := piles[lidx] + piles[lidx+1]
		return max(opt1, opt2)
	}

	opt1 := piles[lidx] - maxDiffForFirstPlayer(piles, lidx+1, cache)
	opt2 := piles[lidx] + piles[lidx+1] - maxDiffForFirstPlayer(piles, lidx+2, cache)
	opt3 := piles[lidx] + piles[lidx+1] + piles[lidx+2] - maxDiffForFirstPlayer(piles, lidx+3, cache)

	l := max(opt1, opt2, opt3)
	cache[lidx] = l

	return l
}

func main() {
	piles1 := []int{1, 2, 3, 7}
	piles2 := []int{1, 2, 3, -9}
	piles3 := []int{1, 2, 3, 6}
	piles4 := []int{-1, -2, -3}
	piles5 := []int{1, 1, 1, 0, 1, 1, 1, 1, 1, 1}
	piles6 := []int{9, -4, 0, 12, -5, -13, 15, 6, -16, 8, 2, 16, 12, -6, 13, 0, -16, -11, 9, -14, 7, -1, 14}
	piles7 := []int{-3, 4, -3, 10, 4, 11, -14, 12, -10, -6, 7, 3, -1, -13, -4, 11, -9, -8, 11, -11, 12, 9, 3}

	fmt.Println(stoneGameIII(piles1))
	fmt.Println(stoneGameIII(piles2))
	fmt.Println(stoneGameIII(piles3))
	fmt.Println(stoneGameIII(piles4))
	fmt.Println(stoneGameIII(piles5))
	fmt.Println(stoneGameIII(piles6))
	fmt.Println(stoneGameIII(piles7))

}

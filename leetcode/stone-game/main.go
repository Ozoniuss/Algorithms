package main

import "fmt"

func stoneGame(piles []int) bool {
	cache := make(map[[2]int]int, 0)
	total := 0
	for i := range piles {
		total += piles[i]
	}
	m := maxStonesForAlice(piles, 0, len(piles), 0, cache)
	return 2*m >= total
}

// assumes alice is at play here
func maxStonesForAlice(piles []int, lidx int, ridx int, currentStones int, cache map[[2]int]int) int {

	// fmt.Println(piles, lidx, ridx, cache, currentStones)
	if ridx < lidx {
		panic("wtf")
	}

	if c, ok := cache[[2]int{lidx, ridx}]; ok {
		return currentStones + c
	}

	if lidx == ridx {
		return currentStones
	}

	if lidx == ridx-1 {
		return currentStones + piles[lidx]
	}

	if lidx == ridx-2 {
		opt1 := currentStones + piles[lidx]
		opt2 := currentStones + piles[ridx-1]

		return max(opt1, opt2)
	}

	opt1 := maxStonesForAlice(piles, lidx+2, ridx, currentStones+piles[lidx], cache)
	opt2 := maxStonesForAlice(piles, lidx+1, ridx-1, currentStones+piles[lidx], cache)

	opt3 := maxStonesForAlice(piles, lidx, ridx-2, currentStones+piles[ridx-1], cache)
	opt4 := maxStonesForAlice(piles, lidx+1, ridx-1, currentStones+piles[ridx-1], cache)

	l := max(min(opt1, opt2), min(opt3, opt4))
	cache[[2]int{lidx, ridx}] = l

	return l

}

func main() {
	piles1 := []int{5, 3, 4, 5}
	piles2 := []int{3, 7, 2, 3}
	piles3 := []int{3, 2, 10, 4}

	fmt.Println(stoneGame(piles1))
	fmt.Println(stoneGame(piles2))
	fmt.Println(stoneGame(piles3))

}

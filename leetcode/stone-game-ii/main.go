package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
}

func stoneGameII(piles []int) int {
	cache := make(map[[2]int]int, 0)
	total := 0
	for _, p := range piles {
		total += p
	}
	return (total + diff(piles, 0, 1, cache)) / 2
}

func diff(piles []int, pidx int, M int, cache map[[2]int]int) int {
	if len(piles) < pidx {
		panic("piles")
	}
	if len(piles) == pidx {
		return 0
	}

	if v, ok := cache[[2]int{pidx, M}]; ok {
		return v
	}

	diffs := make([]int, 0, 2*M)
	for i := 1; i <= 2*M; i++ {
		if pidx+i > len(piles) {
			d := 0
			for j := pidx; j < len(piles); j++ {
				d += piles[j]
			}
			diffs = append(diffs, d)
			break
		}
		d := -diff(piles, pidx+i, max(M, i), cache)
		for j := pidx; j < pidx+i; j++ {
			d += piles[j]
		}
		diffs = append(diffs, d)
	}

	m := slices.Max(diffs)
	cache[[2]int{pidx, M}] = m
	return m
}

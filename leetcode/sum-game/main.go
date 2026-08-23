package main

import "fmt"

// ??
// 1?
// 12

// diff can go up or down by 9

func main() {
	fmt.Println(sumGame("5023"))
	fmt.Println(sumGame("25??"))
	fmt.Println(sumGame("?3295???"))
}

func sumGame(num string) bool {

	if len(num)%2 != 0 {
		panic("len")
	}

	increases := len(num) / 2
	decreases := len(num) / 2
	diff := 0

	for i, c := range num {
		if c == '?' {
			continue
		}
		v := int(c) - int('0')
		if i < len(num)/2 {
			increases -= 1
			diff += v
		} else {
			decreases -= 1
			diff -= v
		}
	}
	return sumGame2(increases, decreases, diff)
}

func sumGame2(increases, decreases, diff int) bool {

	if increases == 0 && decreases == 0 {
		return diff != 0
	}

	// if alice plays the last move, she will always win
	if (increases+decreases)%2 != 0 {
		return true
	}

	// it doesn't really matter in which order we play increases or decreases
	// for one individual
	// if I only have increases left, as alice I will try to get it above 0 if I can, otherwise keep it below 0
	// if I only have decreases left, as alice I will try to get it below 0 if I can, otherwise keep it above 0
	//
	// 1000 increases, 4 decreases
	// -36 496 * 9
	// 1004 increases
	// 502 * 9

	// strategy of alice trying to get everything to the right (diff > 0)
	diffRight := diff
	diffRight += 9 * (increases + 1) / 2 // alice's moves, she will try to use all increases
	diffRight -= 9 * (decreases + 1) / 2 // bob trying to get it to the left

	diffLeft := diff
	diffLeft -= 9 * (decreases + 1) / 2
	diffLeft += 9 * (increases + 1) / 2

	return diffRight > 0 || diffLeft < 0

}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

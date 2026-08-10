package main

import (
	"fmt"
)

func main() {
	fmt.Println(winnerSquareGame(3))
}
func winnerSquareGame(n int) bool {
	SIZE := 100
	// SIZE := 11
	outcomes := make([]bool, SIZE)
	outcomes[0] = false
	outcomes[1] = true

	for i := 2; i*i < SIZE; i++ {
		outcomes[i*i] = true
	}

	for i := 2; i < SIZE; i += 1 {

		if outcomes[i] == true {
			continue
		}

		canwin := false
		// if in any of these I can end up in a losing position,
		// I win here. otherwise, I lose
		for j := 1; j*j < i; j++ {
			canwin = canwin || !outcomes[i-j*j]
		}
		outcomes[i] = canwin
	}
	for i := 1; i < SIZE; i++ {
		if outcomes[i] == false {
			fmt.Println(i, outcomes[i])
		}
	}
	return outcomes[n]
}

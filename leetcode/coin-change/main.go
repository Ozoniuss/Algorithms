package main

import (
	"fmt"
	"slices"
)

// 1,2, 5

// 11

// 3, 4, 6
// 11

// 100000019
// 5 3

// 7999
// 4000, 3, 1

const N = 10001

func coinChangedp(coins []int, amount int) int {
	dp := [N]int{}
	for i := range N {
		dp[i] = -1
	}
	dp[0] = 0
	for _, c := range coins {
		for i := c; i < N; i++ {
			if dp[i-c] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-c] + 1
				} else {
					dp[i] = min(dp[i], dp[i-c]+1)
				}
			}
		}
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	slices.SortFunc(coins, func(a, b int) int {
		return a - b
	})
	fmt.Println(coins)
	return helper(coins, amount, &[N]int{})
}

func helper(coins []int, amount int, cache *[N]int) int {

	if amount == 0 {
		return 0
	}

	// assume coins are sorted in increasing order, cannot make this
	// amount
	if amount < coins[0] {
		(*cache)[amount] = -1
		return -1
	}

	// start from the largest coin
	for i := len(coins) - 1; i > -1; i-- {
		c := coins[i]

		// cannot use this coin, try a smaller one
		if c > amount {
			continue
		}

		// with this coin we got a value that is not 0
		if v := helper(coins, amount-c, cache); v != -1 {
			(*cache)[amount] = 1 + v
			return 1 + v
		}
	}
	(*cache)[amount] = -1
	return -1
}

func main() {
	fmt.Println(coinChangedp([]int{1, 2000000}, 2))
	fmt.Println(coinChangedp([]int{3, 5}, 1019))
	fmt.Println(coinChangedp([]int{55, 13, 11}, 2))
}

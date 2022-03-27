package main

import (
	"fmt"
	"math"
)

// Solution 1 (recursion, less than 2 min to write)
// Complexity: exponential, Th(n) = Th(n-1) + Th(n-2), which approximates to (1/2 + sqrt 5 / 2)^n
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// Solution 2, dynamic programming
func climbStairsDP(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	solutions := []int{1, 2}
	for i := 2; i < n; i++ {
		solutions = append(solutions, solutions[i-1]+solutions[i-2])
	}
	return solutions[n-1]
}

// Solution 3, simply solve x_n = x_{n-1} + x_{n-2}, x_1 = 1, x_2 = 2
func climbStairsFormula(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	eps := (1 + math.Sqrt(5)) / 2
	eps_c := (1 - math.Sqrt(5)) / 2
	return int(math.Round((1 / math.Sqrt(5)) * (math.Pow(eps, float64(n+1)) - math.Pow(eps_c, float64(n-1)))))
}

func main() {
	fmt.Println(climbStairsDP(4))
}

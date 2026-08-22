package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthSmallest([]int{3, 6, 9}, 3))
	fmt.Println(findKthSmallest([]int{2, 5}, 7))
	fmt.Println(findKthSmallest([]int{1, 4}, 5))

}

func findKthSmallest(coins []int, k int) int64 {
	subsets := make([][][]int, len(coins)+1)
	for i := 1; i < len(coins)+1; i++ {
		subsets[i] = generateSubset(coins, 0, i)
	}

	step := 1 << 32
	current := 0

	for step > 0 {
		// we will stop increasing current at exactly the point when the number
		// of combinations with the coins we can make that are strictly smaller
		// than it becomes k (since last step is 1).
		for coinConbinationsLessThanK(coins, current+step, subsets) < k {
			current += step
		}
		step = step / 2
	}
	return int64(current)
}

func coinConbinationsLessThanK(coins []int, k int, subsets [][][]int) int {
	beforek := 0
	for i := 1; i < len(coins)+1; i++ {
		subsetlist := subsets[i]
		for _, s := range subsetlist {
			lc := lcmList(s)
			diff := k / lc
			if k%lc == 0 {
				diff -= 1
			}

			if i%2 != 0 {
				beforek += diff
			} else {
				beforek -= diff
			}
		}
	}
	return beforek
}

func generateSubset(nums []int, pos int, size int) [][]int {
	if size == 0 {
		return [][]int{{}}
	}
	// not enough numbers left, should return
	if pos+size > len(nums) {
		return nil
	}
	all := make([][]int, 0)
	for i := pos; i <= len(nums)-size; i++ {
		el := nums[i]
		subsets := generateSubset(nums, i+1, size-1)
		if len(subsets) != 0 {
			for _, s := range subsets {
				cl := []int{el}
				cl = append(cl, s...)
				all = append(all, cl)
			}
		}
	}
	return all
}

// gcd returns the greatest common divisor of a and b.
// Euclidean algorithm: O(log(min(a, b))).
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// gcdList returns the GCD of all numbers.
// gcd(a, b, c, ...) = gcd(gcd(gcd(a, b), c), ...)
func gcdList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 1; i < len(nums); i++ {
		result = gcd(result, nums[i])

		// 1 is the smallest possible positive GCD,
		// so we can stop early.
		if result == 1 {
			return 1
		}
	}

	if result < 0 {
		result = -result
	}

	return result
}

// lcm returns the least common multiple of a and b.
//
// lcm(a, b) = |a / gcd(a, b) * b|
//
// Divide before multiplying to reduce the chance of overflow.
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := (a / gcd(a, b)) * b

	if result < 0 {
		result = -result
	}

	return result
}

// lcmList returns the LCM of all numbers.
// lcm(a, b, c, ...) = lcm(lcm(lcm(a, b), c), ...)
func lcmList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])

		// If any number is 0, the LCM is 0.
		if result == 0 {
			return 0
		}
	}

	if result < 0 {
		result = -result
	}

	return result
}

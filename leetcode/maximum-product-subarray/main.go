package main

import "fmt"

// 2, 3, -2, 1 -10 4
//                 [4,]
// 1 9 9 9 9 0 9

//   -5      6                    10                 -4
//  1200, -300       60, -240    10, -40       1, -4

//  -2 -4 [......] -> P

// very clever idea i found on lc
func maxProductKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	cp := 1
	maxp := -1 << 31

	for _, n := range nums {
		if n == 0 {
			maxp = max(maxp, 0)
			cp = 1
		} else {
			cp *= n
			maxp = max(cp, maxp)
		}
	}

	for i := len(nums) - 1; i > -1; i-- {
		n := nums[i]
		if n == 0 {
			maxp = max(maxp, 0)
			cp = 1
		} else {
			cp *= n
			maxp = max(cp, maxp)
		}
	}
	return maxp
}

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	// will either have just negative numbers or at least one positive
	// number. so product is definitely going to be greater than 0 or
	// equal to 0

	n := len(nums)
	dp := make([][2]int, n)
	if nums[n-1] == 0 {
		dp[n-1] = [2]int{0, 0}
	} else if nums[n-1] < 0 {
		dp[n-1] = [2]int{-1, -nums[n-1]}
	} else {
		dp[n-1] = [2]int{nums[n-1], -1}
	}

	for i := n - 2; i > -1; i-- {
		prevp := dp[i+1][0]
		prevn := dp[i+1][1]

		if nums[i] == 0 {
			// does not matter, ditch all prevs
			dp[i] = [2]int{0, 0}
		} else if nums[i] > 0 {
			// I only care if I can form a positive sum
			dp[i] = [2]int{prevp * nums[i], prevn * nums[i]}
			if prevp == -1 || prevp == 0 {
				dp[i][0] = nums[i]
			}
			if prevn == -1 || prevn == 0 {
				dp[i][1] = -1
			}
		} else if nums[i] < 0 {
			// I only care if I can form a positive sum
			dp[i] = [2]int{prevn * -nums[i], prevp * -nums[i]}
			if prevn == -1 || prevn == 0 {
				dp[i][0] = -1
			}
			if prevp == -1 || prevp == 0 {
				dp[i][1] = -nums[i]
			}
		}
	}
	fmt.Println(dp)

	maxp := -1 << 31
	for i := 0; i < len(dp); i++ {
		maxp = max(dp[i][0], maxp)
	}

	return maxp
}

// -4     -2
//  8, 4    -1, 2

//	    -5               2       3       -2        4
//	   240, 30         6,48       3, 24       -1, 8           4,-1
//		1, -8  4,1
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-4, -3}))
}

// ----->
// 2 5 3 -4 9 -1 -3 4 8 0

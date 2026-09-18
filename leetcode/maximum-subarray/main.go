package main

import "fmt"

func maxSubArray(nums []int) int {
	ret, _ := maxSubArray2(nums)
	return ret
}

func maxSubArray2(nums []int) (int, []int) {
	curr := 0
	maxs := -1 << 63

	var left int
	var right int
	var cpos int

	for i := range len(nums) {
		curr += nums[i]

		if maxs != max(maxs, curr) {
			maxs = curr
			left = cpos
			right = i + 1
		}

		// change cpos when current sum does not generate a positive
		// number since including it would not help
		if curr < 0 {
			curr = 0
			cpos = i + 1
		}
	}

	return maxs, nums[left:right]
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray2([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray2([]int{5, 4, -1, 7, 8, -1000000000}))
	fmt.Println(maxSubArray2([]int{-1, -1}))
	fmt.Println(maxSubArray2([]int{-1, -2, -3}))
	fmt.Println(maxSubArray2([]int{-3, -2, -1}))
}

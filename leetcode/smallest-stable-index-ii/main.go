package main

import "fmt"

func firstStableIndex(nums []int, k int) int {

	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	maxarr := make([]int, n)
	minarr := make([]int, n)

	maxarr[0] = nums[0]
	minarr[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		maxarr[i] = max(nums[i], maxarr[i-1])
		minarr[n-1-i] = min(nums[n-i-1], minarr[n-i])
	}

	for i := range nums {
		scr := maxarr[i] - minarr[i]
		if scr <= k {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
	fmt.Println(firstStableIndex([]int{3, 2, 1}, 1))
	fmt.Println(firstStableIndex([]int{0}, 0))
}

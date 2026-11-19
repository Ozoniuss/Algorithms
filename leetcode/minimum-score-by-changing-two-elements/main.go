package main

import (
	"slices"
)

func minimizeSum(nums []int) int {
	slices.Sort(nums)

	N := len(nums)
	return min(nums[N-1]-nums[2], nums[N-3]-nums[0], nums[N-2]-nums[1])

}

func main() {

}

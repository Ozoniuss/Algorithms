package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(missingInteger([]int{1, 2, 3, 2, 5}))
	fmt.Println(missingInteger([]int{3, 4, 5, 1, 12, 14, 13}))
	fmt.Println(missingInteger([]int{29, 30, 31, 32, 33, 34, 35, 36, 37}))
	fmt.Println(missingInteger([]int{4, 5, 6, 7, 8, 8, 9, 4, 3, 2, 7}))
}
func missingInteger(nums []int) int {
	seq := []int{}
	s := 0
	totals := 0
	for i := range len(nums) {
		if len(seq) == 0 || nums[i]-seq[len(seq)-1] == 1 {
			seq = append(seq, nums[i])
			s += nums[i]
		} else {
			break
		}
	}

	totals = s

	slices.Sort(nums)
	for _, num := range nums {
		if num == totals {
			totals += 1
		}
	}
	return totals
}

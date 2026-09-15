package main

import (
	"fmt"
)

func minimumDeletions(nums []int) int {
	minpos := 0
	maxpos := 0
	minv := 1<<63 - 1
	maxv := -1 << 63

	for i := range len(nums) {
		if nums[i] > maxv {
			maxv = nums[i]
			maxpos = i
		}
		if nums[i] < minv {
			minv = nums[i]
			minpos = i
		}
	}

	N := len(nums)
	if minpos == maxpos {
		return min(minpos+1, N-minpos)
	}

	v1 := max(minpos+1, maxpos+1)
	v2 := max(N-minpos, N-maxpos)

	v3 := minpos + 1 + N - maxpos
	v4 := maxpos + 1 + N - minpos

	return min(v1, v2, v3, v4)

}

func main() {
	fmt.Println(minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	fmt.Println(minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
	fmt.Println(minimumDeletions([]int{101}))

}

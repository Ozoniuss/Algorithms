package main

import (
	"fmt"
	"slices"
)

// k1, v1
// k2, v2
// k3, v3
// ...
// kn, vn

type element struct {
	indices [2]int
	val     int
}

func fourSum(nums []int, target int) [][]int {
	sols := make(map[[4]int]struct{})
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			left := j + 1
			right := len(nums) - 1
			for left < right {
				total := nums[i] + nums[j] + nums[left] + nums[right]
				if total == target {
					n := []int{nums[i], nums[j], nums[left], nums[right]}
					slices.Sort(n)
					sols[[4]int{n[0], n[1], n[2], n[3]}] = struct{}{}
					left += 1
				} else if total < target {
					left += 1
				} else {
					right -= 1
				}
			}
		}
	}
	// fmt.Print(sols)
	out := make([][]int, 0)
	for s, _ := range sols {
		out = append(out, []int{s[0], s[1], s[2], s[3]})
	}
	return out
}

// 17 15 ... -10

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

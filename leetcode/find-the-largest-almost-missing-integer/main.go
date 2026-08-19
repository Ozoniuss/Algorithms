package main

import "fmt"

func largestIntegerV2(nums []int, k int) int {

	if k == 0 || k > len(nums) {
		panic("invalid k")
	}

	if k == len(nums) {
		m := -1
		for _, v := range nums {
			m = max(m, v)
		}
		return m
	}

	if k == 1 {
		cnt := make(map[int]int)
		for _, v := range nums {
			cnt[v] += 1
		}
		m := -1
		for k, v := range cnt {
			if v == 1 {
				m = max(m, k)
			}
		}
		return m
	}

	if k >= 2 {

		if len(nums) == 2 {
			return max(nums[0], nums[1])
		}

		lCnt := 0
		rCnt := 0
		for _, v := range nums {
			if nums[0] == v {
				lCnt += 1
			}
			if nums[len(nums)-1] == v {
				rCnt += 1
			}
		}
		if lCnt > 1 && rCnt > 1 {
			return -1
		}
		if lCnt == 1 && rCnt == 1 {
			return max(nums[0], nums[len(nums)-1])
		}
		if lCnt == 1 {
			return nums[0]
		}
		if rCnt == 1 {
			return nums[len(nums)-1]
		}
	}

	panic("should never happen")
}

func largestInteger(nums []int, k int) int {
	l := -1
	cnt := make(map[int]int)
	for r := k - 1; r < len(nums); r++ {
		l += 1
		counted := make(map[int]struct{})
		for i := l; i <= r; i++ {
			if _, ok := counted[nums[i]]; !ok {
				cnt[nums[i]] += 1
				counted[nums[i]] = struct{}{}
			}
		}
	}
	m := -1
	for k, v := range cnt {
		if v == 1 {
			m = max(m, k)
		}
	}
	return m
}

func main() {
	fmt.Println(largestIntegerV2([]int{3, 9, 2, 1, 7}, 3))
	fmt.Println(largestIntegerV2([]int{3, 9, 7, 2, 1, 7}, 4))
	fmt.Println(largestIntegerV2([]int{0, 0}, 1))
	fmt.Println(largestIntegerV2([]int{0, 0}, 2))
}

package main

func missingMultiple(nums []int, k int) int {
	cnt := make(map[int]struct{})
	for _, n := range nums {
		cnt[n] = struct{}{}
	}
	for i := k; i < k*(len(nums)+1); i += k {
		if _, ok := cnt[i]; !ok {
			return i
		}
	}
	return k * (len(nums) + 1)
}

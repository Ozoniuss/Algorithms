package main

import "fmt"

func main() {
	fmt.Println(longestContiguousSubsequence([]int{1, 2, 3}))
	fmt.Println(longestContiguousSubsequence([]int{2, 3, 4}))
	fmt.Println(longestContiguousSubsequence([]int{0, 7}))
	fmt.Println(longestContiguousSubsequence([]int{0, 0}))
	fmt.Println(longestContiguousSubsequence([]int{0, 0, 7, 0, 0, 0, 7, 0, 0}))
}

func longestSubsequence(nums []int) int {
	xor := 0
	hasNonzero := false
	for _, n := range nums {
		xor = xor ^ n
		if n != 0 {
			hasNonzero = true
		}
	}

	if xor != 0 {
		return len(nums)
	}
	if !hasNonzero {
		return 0
	}
	return len(nums) - 1
}

// I originally thought it needs to be contiguous
func longestContiguousSubsequence(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor = xor ^ n
	}

	if xor != 0 {
		return len(nums)
	}

	left := 0
	right := 0
	leftxor := xor
	rightxor := xor

	for leftxor == 0 && left < len(nums) {
		leftxor = leftxor ^ nums[left]
		left += 1
	}

	for rightxor == 0 && right < len(nums) {
		rightxor = rightxor ^ nums[len(nums)-right-1]
		right += 1
	}

	return max(len(nums)-left, len(nums)-right)
}

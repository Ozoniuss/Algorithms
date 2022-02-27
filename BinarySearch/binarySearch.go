package main

import "fmt"

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	// create the two pivots
	left_index := 0
	right_index := len(nums)

	for {
		// if left pivot passes right pivot, there is no number
		if left_index > right_index {
			return -1
		}

		midpoint := (left_index + right_index) / 2

		// edge case if searching for a number greater than the last number
		if midpoint == len(nums) {
			return -1
		}

		// all numbers after middle are greater than middle which is greater than target, no need to check those
		if nums[midpoint] > target {
			right_index = midpoint - 1
		}

		// same as above is number in the middle is smaller
		if nums[midpoint] < target {
			left_index = midpoint + 1
		}

		// found the number, return the index
		if nums[midpoint] == target {
			return midpoint
		}
	}

}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

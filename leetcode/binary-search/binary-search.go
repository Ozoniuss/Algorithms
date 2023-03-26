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

// searchForPropertyRight searches for the smallest number such that a property
// holds, iteratively. Assume that the property holds for all numbers, starting
// from one number in the array.
//
// e.g. [x x x x x H H H H H]
func searchForPropertyRight(nums []int, propertyHolds func(int) bool) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums)

	for {

		if left == right {
			return nums[left]
		}

		mid := (left + right) / 2

		// There might be smaller numbers with that property.
		if propertyHolds(nums[mid]) {
			right = mid
			continue
			// Property doesn't hold for this number, so just skip numbers from
			// here.
		} else {
			left = mid + 1
		}
	}
}

// searchForPropertyLeft searches for the greatest number such that a property
// holds, iteratively. Assume that the property holds for all numbers, starting
// from one number in the array.
//
// e.g. [H H H H H H x x x x]
func searchForPropertyLeft(nums []int, propertyHolds func(int) bool) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for {

		if left == right {
			if propertyHolds(nums[left]) {
				return nums[left]
			}
			return -1
		}

		mid := (left + right + 1) / 2

		if propertyHolds(nums[mid]) {
			left = mid
			continue
		} else {
			right = mid - 1
			continue
		}
	}
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(searchForPropertyRight([]int{0, 1, 3, 5, 9, 12}, func(i int) bool { return i >= 1 }))
	fmt.Println(searchForPropertyLeft([]int{0, 1, 3, 5, 9, 12}, func(i int) bool { return i <= -2 }))
}

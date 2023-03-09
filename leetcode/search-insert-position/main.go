package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := (left + right) / 2

		if left == right {
			if nums[left] < target {
				return left + 1
			} else {
				return left
			}
		}

		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid
			continue
		}
		if target > nums[mid] {
			left = mid + 1
			continue
		}
	}
}

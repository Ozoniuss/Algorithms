package main

import "fmt"

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func sortColors(nums []int) {
	low := 0
	high := len(nums) - 1

	i := 0
	for i <= high {
		if i < low {
			i = low
			continue
		}
		if nums[i] == 0 {
			swap(nums, low, i)
			low += 1
		} else if nums[i] == 2 {
			swap(nums, high, i)
			high -= 1
		} else {
			i++
		}
		fmt.Println(i, nums)
	}
}

func main() {
	// nums := []int{2, 0, 2, 1, 1, 0}
	// nums := []int{1, 0}
	nums := []int{1, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}

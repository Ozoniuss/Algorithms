package main

import "fmt"

// Basically, to rephrase the statement, we want the smallest number of the
// sequence 1,2,3,... that is not part of the array.
//
// There are N numbers, so the answer we are looking for is definitely in
// the set {1,2,...,N+1}.
//
// To solve this problem, we will be placing every number i that is in the array
// at position i, as long as i is between 1 and N+1.
// !! This assumes counting from 1 instead of 0, adjust index !!
//
// To do this, we will loop through the array and swap value at the current
// position with the value at position current_value + 1, if the operation
// makes sense (current_value is greater than 1 and smaller than N+1, is
// different than the current_index + 1 and is also different than the
// number we want to swap with)
//
// After all swapping is finished, a linear scan will reveal the first number
// that is different from its index + 1. If all numbers are parsed, it simply
// means that the array is {1,2,...,N} and we want N+1.

func firstMissingPositive(nums []int) int {
	l := len(nums)

	// Perform all swaps.
	for idx := 0; idx < l; idx++ {
		should_swap := true
		for should_swap {
			should_swap = swap(&nums, idx)
		}
	}

	// Identify the missing number
	for idx := range nums {
		if nums[idx] != idx+1 {
			return idx + 1
		}
	}
	return len(nums) + 1
}
func swap(nums *[]int, idx int) bool {

	val := (*nums)[idx]

	// Swapping in this cases is either impossible or doesn't make sense
	// because the value is already at the positiion we want it to be.
	if val < 1 || val > len(*nums) || val == idx+1 {
		return false
	}

	// If the numbers you want to swap are equal, it results in an infinite loop.
	if (*nums)[val-1] == (*nums)[idx] {
		return false
	}

	// Swap the numbers.
	aux := (*nums)[val-1]
	(*nums)[val-1] = val
	(*nums)[idx] = aux
	return true
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

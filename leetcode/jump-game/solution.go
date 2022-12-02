package main

import "fmt"

// The idea here is to start from the last position. Since we can always jump
// one step, we only really care about zeroes in the array, because we must
// jump over them. To jump over k consecutive zeroes, either the first number
// before them must be k+1, the second k+2 and so on. If this doesn't happen
// before the next zero, we keep increasing the counter in case we will be
// able to jump over the batch from some place before in the array.
func canJump(nums []int) bool {

	// Allows to compare the current element with the next element even for
	// the first element of the array.
	last := len(nums) - 1
	nums[last] = 1
	can_jump := true
	jump_size := 0
	for i := last; i >= 0; i-- {

		if nums[i] == 0 && nums[i+1] != 0 {
			// Could not jump over the previous batch, keep increasing jump size
			// because we might jump over the batch from some place before in
			// the array.
			if can_jump == false {
				jump_size++
			} else {
				can_jump = false
				jump_size = 1
			}
		} else
		// Keep going through the batch
		if nums[i] == 0 && nums[i+1] == 0 {
			jump_size++
		} else

		// Out of the batch
		if nums[i] != 0 {
			// If we can jump from a location over a batch we don't really
			// care about other batches.
			if can_jump {
				continue
			}
			jump_size++

			// Check if we can jump over
			if nums[i] >= jump_size {
				can_jump = true
			}
		}

	}
	return can_jump
}

func main() {
	fmt.Println(canJump([]int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}))

}

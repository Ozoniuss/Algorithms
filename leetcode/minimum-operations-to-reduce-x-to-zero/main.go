package main

import (
	"fmt"
	"slices"
)

func minOperations(nums []int, x int) int {

	prefixleft := make([]int, len(nums))
	prefixright := make([]int, len(nums))

	s := 0
	t := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		t += nums[len(nums)-i-1]
		prefixleft[i] = s
		prefixright[i] = t
	}

	minels := 1<<63 - 1
	for i := range len(prefixleft) {
		left := prefixleft[i]
		if left > x {
			break
		}
		if left == x {
			minels = min(minels, i+1)
			break
		}
		j, found := slices.BinarySearch(prefixright, x-left)
		if found && i+j < len(nums)-2 {
			minels = min(minels, i+j+2)
		}
	}

	for i := range len(prefixright) {
		right := prefixright[i]
		if right > x {
			break
		}
		if right == x {
			minels = min(minels, i+1)
			break
		}
		j, found := slices.BinarySearch(prefixleft, x-right)
		if found && i+j < len(nums)-2 {
			minels = min(minels, i+j+2)
		}
	}

	if minels == 1<<63-1 {
		return -1
	}

	return minels
}

func minOperationsOptimized(nums []int, x int) int {

	prefixleft := make([]int, len(nums))
	prefixright := make([]int, len(nums))

	s := 0
	t := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		t += nums[len(nums)-i-1]
		prefixleft[i] = s
		prefixright[i] = t
	}

	if s < x {
		return -1
	}

	minels := 1<<63 - 1
	leftptr := 0
	rightptr := len(nums) - 1
	for leftptr < len(nums) {

		if prefixleft[leftptr] == x {
			minels = min(minels, leftptr+1)
			break
		}

		for rightptr >= 0 && prefixleft[leftptr]+prefixright[rightptr] >= x {
			rightptr--
		}

		rightptr++
		removed := leftptr + rightptr + 2
		if prefixleft[leftptr]+prefixright[rightptr] == x && removed < len(nums) {
			minels = min(minels, removed)
		}
		leftptr++
	}

	leftptr = len(nums) - 1
	rightptr = 0
	for rightptr < len(nums) {

		if prefixright[rightptr] == x {
			minels = min(minels, rightptr+1)
			break
		}

		for leftptr >= 0 && prefixleft[leftptr]+prefixright[rightptr] >= x {
			leftptr--
		}
		leftptr++
		removed := leftptr + rightptr + 2
		if prefixleft[leftptr]+prefixright[rightptr] == x && removed < len(nums) {
			minels = min(minels, removed)
		}
		rightptr++
	}

	if minels == 1<<63-1 {
		return -1
	}

	return minels
}

func main() {
	fmt.Println(minOperationsOptimized([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperationsOptimized([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperationsOptimized([]int{3, 2, 20, 1, 1, 3}, 10))
	fmt.Println(minOperationsOptimized([]int{1, 1}, 3))
	fmt.Println(minOperationsOptimized([]int{1241, 8769, 9151, 3211, 2314, 8007, 3713, 5835, 2176, 8227, 5251, 9229, 904, 1899, 5513, 7878, 8663, 3804, 2685, 3501, 1204, 9742, 2578, 8849, 1120, 4687, 5902, 9929, 6769, 8171, 5150, 1343, 9619, 3973, 3273, 6427, 47, 8701, 2741, 7402, 1412, 2223, 8152, 805, 6726, 9128, 2794, 7137, 6725, 4279, 7200, 5582, 9583, 7443, 6573, 7221, 1423, 4859, 2608, 3772, 7437, 2581, 975, 3893, 9172, 3, 3113, 2978, 9300, 6029, 4958, 229, 4630, 653, 1421, 5512, 5392, 7287, 8643, 4495, 2640, 8047, 7268, 3878, 6010, 8070, 7560, 8931, 76, 6502, 5952, 4871, 5986, 4935, 3015, 8263, 7497, 8153, 384, 1136}, 894887480))
}

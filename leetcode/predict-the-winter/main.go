package main

import "fmt"

func predictTheWinner(nums []int) bool {
	// cache stores basically the best I can do for this configuration
	cache := make(map[[2]int]int)
	return maxPotentialScore(nums, 0, len(nums), 0, cache) >= 0
}

// assume player 1 picks a value here, and we want to figure out if he can
// find some way to have higher than delta
func maxPotentialScore(nums []int, lidx, ridx int, advantage int, cache map[[2]int]int) int {
	// fmt.Println(nums, advantage, cache)
	// see if with my current advantage, I could win. I already know the best
	// possible I can do with this configuration
	if bestPossible, ok := cache[[2]int{lidx, ridx}]; ok {
		return advantage + bestPossible
	}
	if lidx == ridx {
		return advantage
	}
	if ridx-lidx == 1 {
		return advantage + nums[lidx]
	}
	if ridx-lidx == 2 {
		score1 := nums[lidx] + advantage - nums[ridx-1]
		score2 := nums[ridx-1] + advantage - nums[lidx]
		if score1 >= score2 {
			return score1
		} else {
			return score2
		}
	}

	// if I choose left
	opt1 := maxPotentialScore(nums, lidx+2, ridx, advantage+nums[lidx]-nums[lidx+1], cache)
	opt3 := maxPotentialScore(nums, lidx+1, ridx-1, advantage+nums[lidx]-nums[ridx-1], cache)

	// if I choose right
	opt4 := maxPotentialScore(nums, lidx, ridx-2, advantage+nums[ridx-1]-nums[ridx-2], cache)
	opt2 := maxPotentialScore(nums, lidx+1, ridx-1, advantage+nums[ridx-1]-nums[lidx], cache)

	bestScore := max(min(opt1, opt3), min(opt2, opt4))
	cache[[2]int{lidx, ridx}] = bestScore - advantage
	return bestScore
}

func main() {
	nums1 := []int{1, 5, 2}
	nums2 := []int{1, 5, 233, 7}
	nums3 := []int{1, 567, 1, 1, 99, 100}
	nums4 := []int{2, 4, 55, 6, 8}
	nums5 := []int{0, 0, 7, 6, 5, 6, 1}
	nums6 := []int{2, 4, 3, 4, 55, 1, 2, 3, 1, 2, 4, 3, 4, 5, 1}
	fmt.Println(predictTheWinner(nums1))
	fmt.Println(predictTheWinner(nums2))
	fmt.Println(predictTheWinner(nums3))
	fmt.Println(predictTheWinner(nums4))
	fmt.Println(predictTheWinner(nums5))
	fmt.Println(predictTheWinner(nums6))
}

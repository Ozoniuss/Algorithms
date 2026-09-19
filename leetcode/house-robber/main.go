package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = nums[n-1]
	dp[n-2] = max(nums[n-1], nums[n-2])
	for i := n - 2; i > -1; i-- {
		// can either skip this and just use the maximum of the next,
		// or take it and use the maximum of the following next
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}

func main() {

}

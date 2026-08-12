package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	nums := make([]int, 10_000)
	for i := range nums {
		nums[i] = i % 100
	}

	profile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}

	runtime.SetCPUProfileRate(100_000)
	if err := pprof.StartCPUProfile(profile); err != nil {
		runtime.SetCPUProfileRate(0)
		profile.Close()
		log.Fatal(err)
	}

	result := maxSubarrayLength(nums, 17)

	pprof.StopCPUProfile()
	if err := profile.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int, len(nums))
	i := -1
	j := 0
	maxlen := 0
	for j < len(nums) {
		freq[nums[j]]++

		// advance pointer until sequence is valid
		v := freq[nums[j]]
		v2 := v
		for v > k {
			i++
			freq[nums[i]]--
			if nums[i] == nums[j] {
				v--
			}
		}
		freq[nums[j]] = v2

		maxlen = max(maxlen, j-i)
		j += 1
	}
	return maxlen
}

func maxSubarrayLengthV2(nums []int, k int) int {
	freq := make([]int, 1_000_000_000)
	i := -1
	j := 0
	maxlen := 0
	for j < len(nums) {
		freq[nums[j]]++

		// advance pointer until sequence is valid
		for freq[nums[j]] > k {
			i++
			freq[nums[i]] -= 1
		}

		maxlen = max(maxlen, j-i)
		j += 1
	}
	return maxlen
}

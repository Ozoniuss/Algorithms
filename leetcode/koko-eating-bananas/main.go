package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// totalTimeToEat computes the total time to eat all bananas, if you eat a
// certain amount per hour.
func totalTimeToEat(piles []int, amount int) int {
	total := 0
	for _, p := range piles {
		total += p / amount
		if p%amount != 0 {
			total += 1
		}
	}
	return total
}

// max returns the maximum integer of an array, assuming positive elements.
func max(arr []int) int {
	max := math.MinInt
	for _, el := range arr {
		if el > max {
			max = el
		}
	}
	return max
}

func minEatingSpeed(piles []int, h int) int {
	// definitely can't cover all piles
	if len(piles) > h {
		return -1
	}

	left := 1
	right := max(piles)

	for {
		if left == right {
			return left
		}

		mid := (left + right) / 2

		if totalTimeToEat(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
}

// readFile reads the piles and the available hours from a file.
func readFile(in string) (nums []int, hours int) {

	f, _ := os.Open(in)

	s := bufio.NewScanner(f)
	s.Scan()
	numsstr := s.Text()

	s.Scan()
	h := s.Text()

	hours, _ = strconv.Atoi(h)

	for _, numstr := range strings.Split(numsstr, " ") {
		num, _ := strconv.Atoi(numstr)
		nums = append(nums, num)
	}

	return
}

func main() {
	piles, h := readFile("input.txt")
	fmt.Println(minEatingSpeed(piles, h))
}

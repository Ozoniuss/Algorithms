package main

import "fmt"

// Just sum the numbers up
// Other ideas: set or hashmap to store the existing numbers but that
// takes unnecessary space and complexity
// there are n+1 numbers total, starting from 0, which means n*(n+1)/2 sum
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return int(n*(n+1)/2) - sum

}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

//This problem was really easy, ~ 5 seconds to come up with the idea and ~ 3 more minutes to write the solution

package main

import (
	"fmt"
	"strconv"
)

// 4 5 2 3 7

// 4 (5,2,3,7)
// (4,5) (2,3,7)
// (4,5,2) (3,7)
// (4,5,2,3) (7)

// 4 ()

func extract(expression string) ([]int, []byte) {
	last := 0
	nums := []int{}
	ops := []byte{}
	for i := range len(expression) {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			n, err := strconv.Atoi(expression[last:i])
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
			ops = append(ops, expression[i])
			last = i + 1
		}
	}
	n, err := strconv.Atoi(expression[last:])
	if err != nil {
		panic(err)
	}
	nums = append(nums, n)
	return nums, ops
}

func diffWaysToCompute(expression string) []int {
	nums, ops := extract(expression)
	return diffh(nums, ops, 0, len(nums), make(map[[2]int][]int))
}

// 4 5 6 7 8 9
//  + - + - +
// left = 1
// i = 3
// right = 4

func diffh(nums []int, ops []byte, left, right int, cache map[[2]int][]int) []int {
	// fmt.Println(nums, ops)
	// fmt.Println("l,r", left, right)
	if left >= right {
		panic("lol")
	}
	if v, ok := cache[[2]int{left, right}]; ok {
		return v
	}
	if left == right-1 {
		return []int{nums[left:right][0]}
	}
	out := []int{}
	for i := left + 1; i <= right-1; i++ {
		out1 := diffh(nums, ops, left, i, cache)
		out2 := diffh(nums, ops, i, right, cache)
		// fmt.Println("out1, out2", out1, out2)
		op := ops[i-1]
		for _, o1 := range out1 {
			for _, o2 := range out2 {
				switch op {
				case '+':
					out = append(out, o1+o2)
				case '-':
					out = append(out, o1-o2)
				case '*':
					out = append(out, o1*o2)
				default:
					panic("invalid op")
				}
			}
		}
	}
	cache[[2]int{left, right}] = out
	return out
}

func main() {
	fmt.Println(diffWaysToCompute("1123"))
	fmt.Println(diffWaysToCompute("1123+11"))
	fmt.Println(diffWaysToCompute("1123+11*2"))
	fmt.Println(diffWaysToCompute("1123+11*2-4"))
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

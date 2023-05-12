package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}

func main() {

}

package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {

}

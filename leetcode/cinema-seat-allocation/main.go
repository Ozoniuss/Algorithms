package main

import "fmt"

func main() {
	fmt.Println(0b00100000&0b00111100 == 0b00000000)
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	reserved := make(map[int]int)
	for _, seats := range reservedSeats {
		if seats[1] >= 2 && seats[1] <= 9 {
			d := seats[1] - 2
			reserved[seats[0]] += 1 << d
		}
	}

	total := 0
	used := 0
	for i := range reserved {
		used += 1
		canReserveLeft := reserved[i]&0b11110000 == 0
		canReserveRight := reserved[i]&0b00001111 == 0
		canReserveMid := reserved[i]&0b00111100 == 0
		if canReserveLeft && canReserveRight {
			total += 2
		} else if canReserveLeft || canReserveRight || canReserveMid {
			total += 1
		}
	}
	return total + (n-used)*2
}

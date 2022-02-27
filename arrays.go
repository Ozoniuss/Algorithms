package main

import "fmt"

// func arrayStuff(arr []int) {
// 	fmt.Printf("%p\n", &arr[0])
// 	for idx, _ := range arr {
// 		arr[idx] = 2
// 	}
// }

func main() {

	// create a slice of size and cap 5
	a := make([]int, 5)
	// b points to the same underlying array as a
	b := a

	//modify slices through a
	a[0] = 1
	a[1] = 1
	a[2] = 1
	a[3] = 1
	//a[4] = 1
	//a[5] = 1

	// should have the same output
	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("b: ", b, len(b), cap(b))

	fmt.Printf("Location in memory of underlying arrays of a: %p and b: %p\n", &a[0], &b[0])

	//because the array is full, this will create a new underlying array as per documentation of append
	a = append(a, 100)

	a[0] = 100
	//b should still be [1 1 1 1 0], 5, 5
	fmt.Println("Should still be [1 1 1 1 0], 5, 5:", b, len(b), cap(b))

	a[0] = 0

	fmt.Printf("Location in memory of underlying arrays of a: %p and b: %p\n", &a[0], &b[0])

	//b now points to the new a slice
	b = a
	// this will not create a new underlying array, which means that it will also modify b
	a = append(a, 100)

	a[0] = 100
	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("b: ", b, len(b), cap(b))

	fmt.Printf("Location in memory of underlying arrays of a: %p and b: %p\n", &a[0], &b[0])

}

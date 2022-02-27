package main

import "fmt"

func cartezianProductIterativeOptimized(sets [][]int) [][]int {
	// idea based on BFS approach

	// initialize the product
	product := [][]int{{}}

	nrSets := len(sets)

	// we will generate the sets in the following way:
	// take the set at the beginning of the current product
	// find its length to know how many elements were already added, let it be k
	// add all elements from the k+1th set to this set
	// append all the new sets at the end
	for len(product[0]) < nrSets {

		k := len(product[0])

		// go though all the elements in k+1th set
		for _, element := range sets[k] {

			// add all of them to the new product
			product = append(product, append(product[0], element))
		}

		// remove first element from product
		product = product[1:]

	}
	return product
}

// func cartezianProductIterative(sets [][]int) [][]int {
// 	// idea based on BFS approach

// 	// initialize the product
// 	product := [][]int{{}}

// 	// for each set, add all of its elements at the end of all the already existing products
// 	for _, set := range sets {
// 		new_products := [][]int{}
// 		for _, element := range set {
// 			for _, p := range product {
// 				new_products := append(new_products, append(p, element))
// 				//fmt.Println(new_products)
// 			}
// 		}

// 	}
// 	return product
// }

func main() {
	_ = cartezianProductIterativeOptimized([][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}})
	// for _, p := range prods {
	// 	fmt.Println(p)
	// }
	fmt.Println("a")
}

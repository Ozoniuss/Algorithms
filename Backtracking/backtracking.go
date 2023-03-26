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

func cartezianProductIterative(sets [][]int) [][]int {
	// idea based on BFS approach

	// initialize the product
	product := [][]int{{}}

	// for each set, add all of its elements at the end of all the already existing products
	for _, set := range sets {
		new_products := [][]int{}
		for _, element := range set {
			for _, p := range product {
				new_products = append(new_products, append(p, element))
				//fmt.Println(new_products)
			}
		}

		// product will now point to the underlying array that new product points to
		product = new_products

	}
	return product
}

func cartezianProductRecursive(sets [][]int) [][]int {

	if len(sets) == 0 {
		return [][]int{{}}
	}
	product := [][]int{}
	first_set := sets[0]
	other_sets := cartezianProductRecursive(sets[1:])
	for _, element := range first_set {
		for _, set := range other_sets {
			product = append(product, append(set, element))
		}
	}

	return product
}

func bfs(nums []int) {
	q := [][]int{}
	for _, n := range nums {
		q = append(q, []int{n})
	}

	for len(q) > 0 {
		curr := q[0]
		fmt.Println(curr)
		q = q[1:]
		for _, num := range nums {

			currCopy := make([]int, len(curr))
			copy(currCopy, curr)
			currCopy = append(currCopy, num)
			if len(currCopy) <= 3 {
				q = append(q, currCopy)
			}
		}
	}
}

func dfsrec(nums []int, curr []int) {

	if len(curr) == 4 {
		return
	}

	fmt.Println(curr)
	for _, num := range nums {
		currCopy := make([]int, len(curr))
		copy(currCopy, curr)
		currCopy = append(currCopy, num)
		dfsrec(nums, currCopy)

	}
}

func bfsrec(nums []int, q [][]int) {

	if len(q[0]) == 4 {
		return
	}

	curr := q[0]
	fmt.Println(curr)
	q = q[1:]
	for _, num := range nums {
		currCopy := make([]int, len(curr))
		copy(currCopy, curr)
		currCopy = append(currCopy, num)
		q = append(q, currCopy)
	}
	bfsrec(nums, q)
}

func dfs(nums []int) {
	s := [][]int{}
	for _, n := range nums {
		s = append(s, []int{n})
	}

	for len(s) > 0 {
		curr := s[0]
		fmt.Println(curr)
		s = s[1:]
		for i := len(nums) - 1; i >= 0; i-- {
			currCopy := make([]int, len(curr))
			copy(currCopy, curr)

			currCopy = append(currCopy, nums[i])

			if len(currCopy) <= 3 {

				c := [][]int{currCopy}
				s = append(c, s...)
			}
		}
	}
}

func update(n *[]int) {
	*n = append(*n, 6)
	*n = append(*n, 7)
	(*n)[1] = 4
}

func updateMap(m map[int]int) {
	m[2] = 4
}

func updateQ(q []int) {
	q = make([]int, 3)
}

func main() {
	// _ = cartezianProductIterativeOptimized([][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}})

	// prods := cartezianProductRecursive([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	// fmt.Println(prods)
	// for _, p := range prods {
	// 	fmt.Println(p)
	// }

	//dfsrec([]int{1, 2, 3}, []int{})
	dfs([]int{1, 2, 3})

	// q := []int{1, 2, 3, 4, 5}
	// fmt.Println(q)
	// updateQ(q)
	// fmt.Println(q)

}

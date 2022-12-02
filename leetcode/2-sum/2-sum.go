package main

import "fmt"

func twoSum(nums []int, target int) []int {

	//map with numbers as keys and their indexes as values
	numbers := make(map[int][]int)
	//var numbers map[int][]int

	// generate the map
	for idx, num := range nums {
		if _, ok := numbers[num]; !ok {
			numbers[num] = []int{idx} // add first position of num to map

			fmt.Println(numbers[num])
		} else {
			numbers[num] = append(numbers[num], idx) // add the other indexes of the number
		}
	}

	for _, num := range nums {
		if _, ok := numbers[target-num]; ok {

			//if numbers are equal, we must have exactly two
			if target-num == num {
				if len(numbers[num]) == 1 {
					continue
				} else {
					return numbers[num]
				}
			} else {
				return []int{numbers[num][0], numbers[target-num][0]}
			}
		}
	}
	return nil
}

//Not really worth the improvement
func twoSumOnePass(nums []int, target int) []int {
	// same solution as the previous one, do it in one pass

	//map with numbers as keys and their indexes as values
	numbers := make(map[int][]int)
	//var numbers map[int][]int

	// generate the map
	for idx, num := range nums {
		if _, ok := numbers[num]; !ok {
			numbers[num] = []int{idx} // add first position of num to map)

			if _, ok := numbers[target-num]; ok {

				//if numbers are equal, we must have exactly two
				if target-num == num {
					if len(numbers[num]) == 1 {
						continue
					} else {
						return numbers[num]
					}
				} else {
					return []int{numbers[num][0], numbers[target-num][0]}
				}
			}

		} else {
			numbers[num] = append(numbers[num], idx) // add the other indexes of the number
			if _, ok := numbers[target-num]; ok {

				//if numbers are equal, we must have exactly two
				if target-num == num {
					if len(numbers[num]) == 1 {
						continue
					} else {
						return numbers[num]
					}
				} else {
					return []int{numbers[num][0], numbers[target-num][0]}
				}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

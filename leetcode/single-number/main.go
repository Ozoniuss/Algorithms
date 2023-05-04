package main

// singleNumber returns the number that appears only once in the array.
func singleNumberInt32(nums []int) int {
	var repeated = make(map[int32]struct{})

	for _, num := range nums {
		if _, ok := repeated[int32(num)]; ok {
			delete(repeated, int32(num))
		} else {
			repeated[int32(num)] = struct{}{}
		}
	}
	keys := make([]int32, 0, len(repeated))
	for k := range repeated {
		keys = append(keys, k)
	}

	return int(keys[0])
}

// singleNumber returns the number that appears only once in the array.
func singleNumber(nums []int) int {
	var repeated = make(map[int]struct{})

	for _, num := range nums {
		if _, ok := repeated[num]; ok {
			delete(repeated, num)
		} else {
			repeated[num] = struct{}{}
		}
	}

	for k := range repeated {
		return k
	}
	return 0
}

var out = 0

func main() {
	out = singleNumberInt32(allnums)
}

var allnums = []int{
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	1, 1,
	2, 2,
	3,
}

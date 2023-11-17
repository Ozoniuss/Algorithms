package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func isNumberStart(c byte) bool {
	return (c == '-') || (('0' <= c) && (c <= '9'))
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b := bufio.NewReader(f)
	var c byte
	err = nil
	inNumber := false

	sb := &strings.Builder{}
	total := 0

	// assume ascii encoded
	for ; err == nil; c, err = b.ReadByte() {
		if !inNumber && isNumberStart(c) {
			inNumber = true
			sb.WriteByte(c)
		} else if inNumber && isNumberStart(c) {
			sb.WriteByte(c)
		} else if inNumber && !isNumberStart(c) {
			inNumber = false
			number := sb.String()
			sb.Reset()

			numberValue, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			total += numberValue
		}
	}
	fmt.Println(total)
}

func computeNodeSum(n any) int {
	switch n.(type) {
	case int:
		return computeNodeSumInteger(n.(int))
	case float64:
		return computeNodeSumInteger(int(n.(float64)))
	case string:
		return computeNodeSumString(n.(string))
	case map[string]any:
		return computeNodeSumMap(n.(map[string]any))
	case []any:
		return computeNodeSumArray(n.([]any))
	default:
		fmt.Println(n)
		fmt.Println(reflect.TypeOf(n))
		panic("wtf")
	}
	return 0
}

func computeNodeSumInteger(n int) int {
	return n
}
func computeNodeSumString(n string) int {
	return 0
}
func computeNodeSumArray(n []any) int {
	s := 0
	for _, el := range n {
		s += computeNodeSum(el)
	}
	return s
}

func computeNodeSumMap(n map[string]any) int {
	s := 0
	for _, val := range n {
		// if val == "red" {
		// 	return 0
		// }
		s += computeNodeSum(val)
	}
	return s
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var alldata map[string]any
	dec := json.NewDecoder(f)
	err = dec.Decode(&alldata)
	if err != nil {
		panic(err)
	}

	fmt.Println(computeNodeSum(alldata))
}

func main() {
	part2()
}

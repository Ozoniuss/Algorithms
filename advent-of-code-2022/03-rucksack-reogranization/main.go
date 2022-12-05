package main

import (
	"bufio"
	"fmt"
	"os"
)

// toggleBit sets the bit at the specified position to 1. Note that allBytes is
// normally passed via copy, so we must provide a pointer in this case.
func toggleBit(position byte, allBytes *[7]byte) {

	// Letters only have priority between 1 and 52.
	if position < 1 || position > 52 {
		return
	}

	// Consider the byte array as follows
	// --------|--------|--------|--------|--------|--------|--------|
	// <section>
	//
	// section:  _______ ________ ________ ________ ________ ________ ________
	//           <location>
	// section represents the position of the byte in the array, location
	// represents the position of the bit in the section.

	section := (position - 1) / 8
	location := (8 - position) % 8

	allBytes[section] = allBytes[section] | (1 << location)
}

// convertToBitArray takes a list of items (characters) and converts them to a
// bit array, such that the bit numbered "i" (where i is between 1 and 52)
// represents whether the character with that priority is in items or not.
func convertToBitArray(items string) [7]byte {

	// Initialize the bit array with zeroes.
	var bitArray [7]byte

	for _, character := range items {
		priority := findItemPriority(character)

		// priority is also the location of the item in the bit array
		toggleBit(priority, &bitArray)
	}

	return bitArray
}

// findCommonItem takes all letters of a single line representing the items in
// both compartiments of a rucksack and finds the common item.
func findCommonItemPriority(items string) byte {

	// Since 7 bytes are 56 bits, we don't need more space to store.
	var leftSideBits [7]byte
	var rightSideBits [7]byte
	var commonItemsBits [7]byte

	// The point where the strings can be split in half. It is assumed that all
	// strings have even length.
	midpoint := len(items) / 2

	// The two parts of the string
	leftSide := items[:midpoint]
	rightSide := items[midpoint:]

	leftSideBits = convertToBitArray(leftSide)
	rightSideBits = convertToBitArray(rightSide)

	for i := 0; i < len(leftSideBits); i++ {
		commonItemsBits[i] = leftSideBits[i] & rightSideBits[i]
	}

	return findItemPriorityFromBitIntersection(commonItemsBits)

}

func findCommonItemsPriority(items [3]string) byte {
	var a [7]byte
	var b [7]byte
	var c [7]byte

	var x [7]byte

	a = convertToBitArray(items[0])
	b = convertToBitArray(items[1])
	c = convertToBitArray(items[2])

	for i := 0; i < len(x); i++ {
		x[i] = a[i] & b[i] & c[i]
	}

	return findItemPriorityFromBitIntersection(x)
}

func findItemPriority(item rune) byte {

	itemChar := byte(item)

	// small letter
	if 97 <= itemChar && itemChar <= 122 {
		return itemChar - 96
	}

	// big letter
	if 65 <= itemChar && itemChar <= 90 {
		return itemChar - 38
	}

	return itemChar
}

func findItemPriorityFromBitIntersection(intersection [7]byte) byte {

	for i := 0; i < len(intersection); i++ {
		// We are only interested in the non-null byte of the set
		if intersection[i] > 0 {

			offset := byte(9)
			val := intersection[i]

			for val > 0 {
				val = val >> 1
				offset -= 1
			}
			return offset + byte(i*8)
		}
	}
	return 0
}

func findPrioritySum() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Scanner that scans lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		items := scanner.Text()
		priority := findCommonItemPriority(items)
		sum += int(priority)
	}

	return sum, nil
}

func findPrioritySumGroups() (int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Scanner that scans lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	groupCounter := 0
	var groupItems [3]string

	for scanner.Scan() {
		items := scanner.Text()
		if groupCounter < 3 {
			groupItems[groupCounter] = items
			groupCounter += 1
		}
		if groupCounter == 3 {
			sum += int(findCommonItemsPriority(groupItems))
			groupCounter = 0
		}

	}

	return sum, nil
}

func main() {
	sum, err := findPrioritySum()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The sum of priorities of the common items is %d\n", sum)

	sumGroup, err := findPrioritySumGroups()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The sum of priorities of the common items is %d\n", sumGroup)
}

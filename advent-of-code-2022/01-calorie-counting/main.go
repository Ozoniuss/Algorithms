package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getCaloriesOfTopElf() (maxCalories int32, maxLine int32, err error) {

	/* Part 1 */

	// Loop through the entire list and use a counter to determine the number
	// of calories carried by every elf.

	// Keep the max in a variable, which is reset if the counter surpasses it.

	var calorieCounter int32 = 0
	var lineNumber int32 = 0

	maxCalories = 0
	maxLine = 0

	f, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	// Initialize a new scanner to read the file line by line.
	scanner := bufio.NewScanner(f)

	// ScanLines uses newline as the line splitter (should be the default
	// splitter set by NewScanner).
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineNumber += 1
		value := scanner.Text()

		// All food calories of the current Elf had been read.
		if value == "" {
			if calorieCounter > maxCalories {
				maxCalories = calorieCounter
				maxLine = lineNumber
			}
			calorieCounter = 0
		} else {
			// Still reading from current elf.
			calories, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return 0, 0, err
			}

			// Increase the calories counter for the current elf.
			calorieCounter += int32(calories)
		}
	}

	// Address EOF special case, calorieCounter stores the value of the last
	// elf
	if calorieCounter > maxCalories {
		maxCalories = calorieCounter
		maxLine = lineNumber
	}

	return maxCalories, maxLine, nil
}

/* Part 2 */

// addToTopThree is a helper method that adds the current elf to the top three
// list, according to its calorie count. If the elf doesn't have enough calories
// to fit in the list, it is automatically not added.
func addToTopThree(top map[int32]int32, calorieCount int32, lineLocation int32) {

	// If there are not enough elements in the map yet, it means that there
	// haven't been enough elfs collected, so just add the value.
	if len(top) < 3 {
		top[calorieCount] = lineLocation
		return
	}

	// Max int32 value
	min := int32(1<<31 - 1)

	for ccount := range top {
		if ccount < min {
			min = ccount
		}
	}

	// The current calorie count is greater than the smallest value in the top.
	// Replace that smallest value (not another one) with the current calorie
	// count
	if calorieCount > min {
		delete(top, min)
		top[calorieCount] = lineLocation
		return
	}
}

func getCaloriesOfTopThreeElfs() (map[int32]int32, error) {

	var calorieCounter int32 = 0
	var lineNumber int32 = 0

	top := make(map[int32]int32)

	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Initialize a new scanner to read the file line by line.
	scanner := bufio.NewScanner(f)

	// ScanLines uses newline as the line splitter (should be the default
	// splitter set by NewScanner).
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineNumber += 1
		value := scanner.Text()

		// All food calories of the current Elf had been read.
		if value == "" {
			addToTopThree(top, calorieCounter, lineNumber)
			calorieCounter = 0
		} else {
			// Still reading from current elf.
			calories, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, err
			}

			// Increase the calories counter for the current elf.
			calorieCounter += int32(calories)
		}
	}

	// Address EOF special case, calorieCounter stores the value of the last
	// elf
	if calorieCounter > 0 {
		addToTopThree(top, calorieCounter, lineNumber)
	}

	return top, nil
}

func main() {

	/* Part 1 */

	maxCalories, maxLine, err := getCaloriesOfTopElf()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The elf carrying most calories is located at %d and carries %d calories.", maxLine, maxCalories)

	top, err := getCaloriesOfTopThreeElfs()
	if err != nil {
		panic(err)
	}

	sum := int32(0)

	fmt.Println("The top 3 elfs carrying most calories are the following:")
	for cal, line := range top {
		fmt.Printf("Located at line %d which carries %d calories;\n", line, cal)
		sum += cal
	}

	fmt.Printf("The sum of all calories is %d", sum)

}

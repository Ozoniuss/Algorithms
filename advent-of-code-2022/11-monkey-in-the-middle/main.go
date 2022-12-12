package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// This problem has a stupid catch. If worry levels are no longer divided by 3,
// the worry levels start growing exponentially and for some reason fuck up
// the division. In order to keep the division under control, I always divide
// the worry levels by the least common multiple of all monkey divisors. This
// ensures that the division operations always produce the same result, while
// keeping the size of the numbers under control.
//
// Note that the operand of each operation is divided by this divisor, not only
// the result.
//
// Obviously I can make this nicer but it's 2 a.m. already and I got sick of
// this problem >:(.
const DIVISOR = 5 * 17 * 2 * 7 * 3 * 11 * 13 * 19

type monkey struct {

	// Worry levels of items held by monkey.
	items []int

	// The operation performed by the monkey on each item.
	operation func(int) int

	// The test performed by the monkey to determine which monkey receives the
	// item next, based on the worry level.
	test func(int) int
}

// readMonkey processes a monkey as a string and returns the monkey data as well
// as the number of the monkey, which should be the position of the monkey in
// the monkey array.
func readMonkey(lines []string) (int, monkey) {

	reNumber := regexp.MustCompile(`[0-9]+`)
	reItems := regexp.MustCompile(`([0-9]+,\s)*[0-9]+`)
	reOperation := regexp.MustCompile(`old.*`)

	monkeyNo, _ := strconv.Atoi(reNumber.FindString(lines[0]))

	divisibleBy, _ := strconv.Atoi(reNumber.FindString(lines[3]))
	throwsIfTrue, _ := strconv.Atoi(reNumber.FindString(lines[4]))
	throwsIfFalse, _ := strconv.Atoi(reNumber.FindString(lines[5]))

	test := func(worryLevel int) int {
		if worryLevel%divisibleBy == 0 {
			return throwsIfTrue
		} else {
			return throwsIfFalse
		}
	}

	items := []int{}
	itemsStrings := strings.Split(reItems.FindString(lines[1]), ", ")

	for _, item := range itemsStrings {
		value, _ := strconv.Atoi(item)
		items = append(items, value)
	}

	operationArgs := strings.Split(reOperation.FindString(lines[2]), " ")

	operation := func(old int) int {

		var operand int
		operandString := operationArgs[2]
		if operandString == "old" {
			operand = old
		} else {
			operand, _ = strconv.Atoi(operandString)
		}

		if operationArgs[1] == "*" {
			return (old % DIVISOR) * (operand % DIVISOR)
		} else if operationArgs[1] == "+" {
			return (old % DIVISOR) + (operand % DIVISOR)
		}

		return old
	}

	return monkeyNo, monkey{
		items:     items,
		operation: operation,
		test:      test,
	}
}

// turn does all operations performed by a single monkey in a round, while also
// returning the total number of items inspected by the monkeys.
func turn(monkeyNo int, monkeys *[]monkey, tooWorried bool) int {

	operations := len((*monkeys)[monkeyNo].items)

	for len((*monkeys)[monkeyNo].items) > 0 {
		worry := (*monkeys)[monkeyNo].operation((*monkeys)[monkeyNo].items[0])
		if !tooWorried {
			worry = worry / 3
		}
		throwsTo := (*monkeys)[monkeyNo].test(worry)
		(*monkeys)[throwsTo].items = append((*monkeys)[throwsTo].items, worry)

		// fmt.Printf("Monkey %d throws item with worry %d to monkey %d\n", monkeyNo, worry, throwsTo)

		// Remove the item to simulate throwing.
		(*monkeys)[monkeyNo].items = (*monkeys)[monkeyNo].items[1:]
	}

	return operations
}

// round performs all operations by all monkeys in a round, including adding
// the number of inspected items of each monkey to the inspected map.
func round(monkeys *[]monkey, inspected *map[int]int, tooWorried bool) {
	for i := 0; i < len(*monkeys); i++ {
		inspectedItems := turn(i, monkeys, tooWorried)
		(*inspected)[i] += inspectedItems
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	monkeyBuf := []string{}

	NO_MONKEYS := 8

	monkeys := make([]monkey, NO_MONKEYS, NO_MONKEYS)
	inspected := make(map[int]int, NO_MONKEYS)

	// Set the number of inspected items by each monkey to 0.
	for i := 0; i < len(monkeys); i++ {
		inspected[i] = 0
	}

	// Read all monkeys.
	// Note that this requires to add newlines at the end of the input file.
	// This isn't nice but the input parsing sucks already so sue me.
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			monkeyBuf = append(monkeyBuf, text)
		} else {
			no, monkey := readMonkey(monkeyBuf)
			monkeys[no] = monkey
			monkeyBuf = []string{}
		}
	}

	worried := true

	fmt.Println(monkeys)

	if !worried {
		//Perform monkey bussiness. Not to worried yet.
		for i := 0; i < 20; i++ {
			round(&monkeys, &inspected, worried)
		}
		for m, i := range inspected {
			fmt.Printf("Monkey %d inspected %d items\n", m, i)
		}
	} else {
		// Way to worried, chill out dude, smoke smth
		for i := 0; i < 10000; i++ {
			round(&monkeys, &inspected, worried)
		}
		for m, i := range inspected {
			fmt.Printf("Monkey %d inspected %d items\n", m, i)
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func processLine(line string, monkeysmap map[string]monkey) {
	monkeyregex := regexp.MustCompile("[a-z]{4}")
	operationregex := regexp.MustCompile(`[+\-*/]`)
	numberregex := regexp.MustCompile("[0-9]+")

	monkeys := monkeyregex.FindAllString(line, -1)
	operation := operationregex.FindString(line)
	number := numberregex.FindString(line)

	if len(monkeys) == 1 {
		yells, _ := strconv.Atoi(number)
		monkeysmap[monkeys[0]] = monkey{
			name:  monkeys[0],
			yells: yells,
		}
	} else {
		monkeysmap[monkeys[0]] = monkey{
			name:      monkeys[0],
			operation: operation,
			waitsFor:  [2]string{monkeys[1], monkeys[2]},
		}
	}
}

// monkey holds the information associated with a monkey, that is, its name,
// the two monkeys it waits for and the operation, or the number it yells.
// It's possible to hold all of these three if the monkeys yelled the numbers
// already.
type monkey struct {
	name      string
	waitsFor  [2]string
	operation string
	yells     int
}

func findYells(m string, monkeys map[string]monkey) int {

	monke := monkeys[m]
	if monke.operation == "" {
		return monke.yells
	} else {
		switch monke.operation {
		case "+":
			return findYells(monke.waitsFor[0], monkeys) + findYells(monke.waitsFor[1], monkeys)
		case "-":
			return findYells(monke.waitsFor[0], monkeys) - findYells(monke.waitsFor[1], monkeys)
		case "/":
			return findYells(monke.waitsFor[0], monkeys) / findYells(monke.waitsFor[1], monkeys)
		case "*":
			return findYells(monke.waitsFor[0], monkeys) * findYells(monke.waitsFor[1], monkeys)
		default:
			return 0
		}
	}
}

func hasHumn(m string, monkeys map[string]monkey) bool {
	monke := monkeys[m]
	if monke.operation == "" {
		return monke.name == "humn"
	} else {
		return hasHumn(monke.waitsFor[0], monkeys) || hasHumn(monke.waitsFor[1], monkeys)
	}
}

// finds the mathematical equation which computes the number yelled by the
// monkey.
func findExpression(m string, monkeys map[string]monkey) string {

	monke := monkeys[m]
	if monke.operation == "" {
		if monke.name == "humn" {
			return "x"
		}
		return fmt.Sprint(monke.yells)

	} else {
		return fmt.Sprintf("(%s%s%s)", findExpression(monke.waitsFor[0], monkeys), monke.operation, findExpression(monke.waitsFor[1], monkeys))
	}
}

// finds the mathematical equation which computes the number yelled by the
// monkey, simplifying where possible.
func findSimplifiedExpression(m string, monkeys map[string]monkey) string {
	monke := monkeys[m]
	if monke.operation == "" {
		if monke.name == "humn" {
			return "x"
		}
		return fmt.Sprint(monke.yells)

	} else {
		if hasHumn(monke.waitsFor[1], monkeys) {
			return fmt.Sprintf("(%d%s%s)", findYells(monke.waitsFor[0], monkeys), monke.operation, findSimplifiedExpression(monke.waitsFor[1], monkeys))
		} else if hasHumn(monke.waitsFor[0], monkeys) {
			return fmt.Sprintf("(%s%s%d)", findSimplifiedExpression(monke.waitsFor[0], monkeys), monke.operation, findYells(monke.waitsFor[1], monkeys))
		} else {
			return fmt.Sprint(findYells(monke.name, monkeys))
		}
	}
}

func computeRecursively(m string, monkeys map[string]monkey, currentValue int) int {
	monke := monkeys[m]
	if monke.operation == "" {
		if monke.name == "humn" {
			return currentValue
		}
		return currentValue
	} else {
		if hasHumn(monke.waitsFor[1], monkeys) {
			value := findYells(monke.waitsFor[0], monkeys)
			switch monke.operation {
			case "+":
				return computeRecursively(monke.waitsFor[1], monkeys, currentValue-value)
			case "-":
				return computeRecursively(monke.waitsFor[1], monkeys, value-currentValue)
			case "/":
				return computeRecursively(monke.waitsFor[1], monkeys, value/currentValue)
			case "*":
				return computeRecursively(monke.waitsFor[1], monkeys, currentValue/value)
			default:
				return 0
			}
		} else if hasHumn(monke.waitsFor[0], monkeys) {
			value := findYells(monke.waitsFor[1], monkeys)
			switch monke.operation {
			case "+":
				return computeRecursively(monke.waitsFor[0], monkeys, currentValue-value)
			case "-":
				return computeRecursively(monke.waitsFor[0], monkeys, currentValue+value)
			case "/":
				return computeRecursively(monke.waitsFor[0], monkeys, currentValue*value)
			case "*":
				return computeRecursively(monke.waitsFor[0], monkeys, currentValue/value)
			default:
				return 0
			}
		}
	}
	return 0
}

func findEqual(monkeys map[string]monkey) {
	root := monkeys["root"]
	// fmt.Println(findExpression(root.waitsFor[0], monkeys))
	// fmt.Println(findExpression(root.waitsFor[1], monkeys))

	value := 0
	//expression := ""
	if hasHumn(root.waitsFor[0], monkeys) {
		value = findYells(root.waitsFor[1], monkeys)
		//expression = findSimplifiedExpression(root.waitsFor[0], monkeys)
		fmt.Println(computeRecursively(root.waitsFor[0], monkeys, value))
	} else {
		value = findYells(root.waitsFor[0], monkeys)
		//expression = findSimplifiedExpression(root.waitsFor[1], monkeys)
		fmt.Println(computeRecursively(root.waitsFor[1], monkeys, value))
	}

}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	monkeys := make(map[string]monkey)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		processLine(scanner.Text(), monkeys)
	}

	root := findYells("root", monkeys)
	fmt.Println(root)

	findEqual(monkeys)

}

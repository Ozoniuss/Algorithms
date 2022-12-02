package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape int8

// Shape values
const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

/* Part 1 */

func computeRoundScoreWinning(enemy, player string) int8 {

	// Shapes associated with the enemy
	var enemyMap = map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	// Shapes associated with the player
	var playerMap = map[string]Shape{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	// Same shape
	if enemyMap[enemy] == playerMap[player] {
		return int8(playerMap[player]) + 3

		// enemy wins
	} else if (enemyMap[enemy] == Rock && playerMap[player] == Scissors) ||
		(enemyMap[enemy] == Scissors && playerMap[player] == Paper) ||
		(enemyMap[enemy] == Paper && playerMap[player] == Rock) {
		return 0 + int8(playerMap[player])

		// player wins
	} else {
		return 6 + int8(playerMap[player])
	}

}

func computeScore(computeMethod func(string, string) int8) (int, error) {

	f, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	playerScore := 0

	// Scanner that scans lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value := scanner.Text()
		parts := strings.Split(value, " ")

		enemy := parts[0]
		player := parts[1]

		score := computeMethod(enemy, player)
		playerScore += int(score)
	}

	return playerScore, nil

}

/* Part 2 helper */

func computeRoundScoreCorrect(enemy, player string) int8 {

	// What beats what
	beats := map[Shape]Shape{
		Rock:     Scissors,
		Scissors: Paper,
		Paper:    Rock,
	}

	losesTo := map[Shape]Shape{
		Rock:     Paper,
		Scissors: Rock,
		Paper:    Scissors,
	}

	// Shapes associated with the enemy
	var enemyMap = map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	// has to lose
	if player == "X" {
		return 0 + int8(beats[enemyMap[enemy]])
	} else if player == "Y" {
		return 3 + int8(enemyMap[enemy])
	} else {
		return 6 + int8(losesTo[enemyMap[enemy]])
	}
}

func main() {
	scoreWinning, err := computeScore(computeRoundScoreWinning)
	if err != nil {
		panic(err)
	}

	scoreCorrect, err := computeScore(computeRoundScoreCorrect)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Winning strategy got score %d\n", scoreWinning)
	fmt.Printf("Correct strategy got score %d\n", scoreCorrect)
}

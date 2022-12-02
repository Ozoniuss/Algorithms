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

type Strategy interface {
	// This method takes the letters representing the action of each player, and
	// returns the round score, based on how the functions implementing this
	// interface interpret those letters.
	computeRoundScore(enemy, player string) int8
}

type CorrectStrategy struct{}

func (s *CorrectStrategy) computeRoundScore(enemy, player string) int8 {
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

type WinningStrategy struct{}

func (s *WinningStrategy) computeRoundScore(enemy, player string) int8 {
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

func computeScore(strategy Strategy) (int, error) {

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

		score := strategy.computeRoundScore(enemy, player)
		playerScore += int(score)
	}

	return playerScore, nil

}

func main() {
	strategyWinning := WinningStrategy{}
	strategyCorrect := CorrectStrategy{}

	scoreWinning, _ := computeScore(&strategyWinning)
	scoreCorrect, _ := computeScore(&strategyCorrect)

	fmt.Printf("Winning strategy got score %d\n", scoreWinning)
	fmt.Printf("Correct strategy got score %d\n", scoreCorrect)
}

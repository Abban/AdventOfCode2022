package main

import (
	"fmt"
	"os"
	"strings"
)

const Rock string = "A"
const Paper string = "B"
const Scissors string = "C"

func main() {
	dir, _ := os.Getwd()
	data, _ := os.ReadFile(dir + "/src/day-02/input.txt")
	dataArray := strings.Split(string(data), "\n")
	score := 0

	for i := 0; i < len(dataArray); i++ {
		if dataArray[i] == "" {
			continue
		}
		hands := strings.Split(dataArray[i], " ")
		hands[1] = convertPlayerHand(hands[1])

		score += getScore(hands[0], hands[1])
	}

	fmt.Println("SCORE: ", score)
}

func getScore(opponentHand string, playerHand string) int {
	playerScore := getPlayerScore(playerHand)

	// Draw
	if opponentHand == playerHand {
		return 3 + playerScore
	}

	// Win
	if opponentHand == Rock && playerHand == Paper {
		return 6 + playerScore
	}

	// Win
	if opponentHand == Paper && playerHand == Scissors {
		return 6 + playerScore
	}

	// Win
	if opponentHand == Scissors && playerHand == Rock {
		return 6 + playerScore
	}

	// Lose
	return playerScore
}

func convertPlayerHand(playerHand string) string {
	switch playerHand {
	case "X":
		return Rock
	case "Y":
		return Paper
	default:
		return Scissors
	}
}

func getPlayerScore(playerHand string) int {
	switch playerHand {
	case Rock:
		return 1
	case Paper:
		return 2
	default:
		return 3
	}
}

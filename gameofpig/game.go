package main

import (
	"fmt"
)

const goalScore = 100

type Player struct {
	name            string
	minStrategy     int
	maxStrategy     int
	currentStrategy int
	dice            Dice
}

func constantStrategy(strategy1 *Player, strategy2 *Player) (int, int, int, int) {
	player1TotalWins := 0
	for i := 0; i < 10; i++ {
		currentTurn := strategy1
		play1TotalScore := 0
		play2TotalScore := 0

		for {
			play1TotalScore += strategyOutcome(strategy1)
			if play1TotalScore >= goalScore {
				break
			}
			currentTurn = strategy2
			play2TotalScore += strategyOutcome(strategy2)
			if play2TotalScore >= goalScore {
				break
			}
			currentTurn = strategy1
		}
		if strategy1 == currentTurn {
			player1TotalWins += 1
		}
	}
	return strategy1.currentStrategy, strategy2.currentStrategy, player1TotalWins, 10
}

func printOutcome(player1Strategy int, player2Strategy int, player1TotalWins int, totalGames int) {
	player2TotalWins := totalGames - player1TotalWins
	player1WinPercentage := float32(player1TotalWins) * 100 / float32(totalGames)
	player2WinPercentage := float32(player2TotalWins) * 100 / float32(totalGames)

	fmt.Printf("Holding at %d vs Holding at %d wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		player1Strategy, player2Strategy,
		player1TotalWins, totalGames, player1WinPercentage,
		player2TotalWins, totalGames, player2WinPercentage)
}

func printMultipleGamesAtOnce(player1Strategy int, player1TotalWins int, totalGames int) {
	player2TotalWins := totalGames - player1TotalWins
	player1WinPercentage := float32(player1TotalWins) * 100 / float32(totalGames)
	player2WinPercentage := float32(player2TotalWins) * 100 / float32(totalGames)

	fmt.Printf("Result: Wins, losses staying at k =  %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
		player1Strategy,
		player1TotalWins, totalGames, player1WinPercentage,
		player2TotalWins, totalGames, player2WinPercentage)
}

func constantAndVariableStrategy(player1 *Player, player2 *Player) {
	var variableStrategy *Player
	var fixedStrategy *Player
	if player1.maxStrategy == player1.minStrategy {
		variableStrategy = player2
		fixedStrategy = player1
	} else {
		variableStrategy = player1
		fixedStrategy = player2
	}
	fixedStrategy.currentStrategy = fixedStrategy.minStrategy
	for i := variableStrategy.minStrategy; i <= variableStrategy.maxStrategy; i++ {
		if i == fixedStrategy.currentStrategy {
			continue
		}
		variableStrategy.currentStrategy = i
		printOutcome(constantStrategy(fixedStrategy, variableStrategy))
	}
}

func variableAndVariableStrategy(player1 *Player, player2 *Player) {
	for i := player1.minStrategy; i <= player1.maxStrategy; i++ {
		counter := 0
		player1WinsSingleStrategy := 0
		player1.currentStrategy = i
		for j := player2.minStrategy; j <= player2.maxStrategy; j++ {
			if i == j {
				continue
			}
			player2.currentStrategy = j
			_, _, player1TotalWins, _ := constantStrategy(player1, player2)
			player1WinsSingleStrategy += player1TotalWins
			counter++
		}
		printMultipleGamesAtOnce(i, player1WinsSingleStrategy, (counter)*10)
	}
}

func strategyOutcome(player *Player) int {
	currentScore := 0
	for {
		number := player.dice.Roll()
		if number == 1 {
			return 0
		}
		if currentScore+number >= player.currentStrategy {
			return currentScore + number
		}
		currentScore += number
	}
}

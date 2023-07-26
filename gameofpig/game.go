package main

import "fmt"

const goalScore = 100

type Player struct {
	name     string
	strategy int
	dice     Dice
}

func constantStrategy(player1Strategy int, player2Strategy int) {
	player1TotalWins := 0
	player1 := &Player{
		name:     "player1",
		strategy: player1Strategy,
		dice:     Dice{},
	}
	player2 := &Player{
		name:     "player2",
		strategy: player1Strategy,
		dice:     Dice{},
	}

	for i := 0; i < 10; i++ {
		currentTurn := player1
		play1TotalScore := 0
		play2TotalScore := 0

		for {
			play1TotalScore += strategyOutcome(player1Strategy, player1)
			if play1TotalScore >= goalScore {
				break
			}
			currentTurn = player2
			play2TotalScore += strategyOutcome(player2Strategy, player2)
			if play2TotalScore >= goalScore {
				break
			}
			currentTurn = player1
		}
		if player1 == currentTurn {
			player1TotalWins += 1
		}
	}
	printOutcome(player1Strategy, player2Strategy, player1TotalWins, 10)
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

func strategyOutcome(strategy int, player *Player) int {
	currentScore := 0
	for {
		number := player.dice.Roll()
		if number == 1 {
			return 0
		}
		if currentScore+number >= strategy {
			return currentScore + number
		}
		currentScore += number
	}
}

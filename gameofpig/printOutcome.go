package main

import "fmt"

func PrintOutcome(player1Strategy int, player2Strategy int, player1TotalWins int, totalGames int) {
	player2TotalWins := totalGames - player1TotalWins
	player1WinPercentage := float32(player1TotalWins) * 100 / float32(totalGames)
	player2WinPercentage := float32(player2TotalWins) * 100 / float32(totalGames)

	fmt.Printf("Holding at %d vs Holding at %d wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		player1Strategy, player2Strategy,
		player1TotalWins, totalGames, player1WinPercentage,
		player2TotalWins, totalGames, player2WinPercentage)
}

func PrintMultipleGamesAtOnce(player1Strategy int, player1TotalWins int, totalGames int) {
	player2TotalWins := totalGames - player1TotalWins
	player1WinPercentage := float32(player1TotalWins) * 100 / float32(totalGames)
	player2WinPercentage := float32(player2TotalWins) * 100 / float32(totalGames)

	fmt.Printf("Result: Wins, losses staying at k =  %d: %d/%d (%.1f%%), %d/%d (%.1f%%)\n",
		player1Strategy,
		player1TotalWins, totalGames, player1WinPercentage,
		player2TotalWins, totalGames, player2WinPercentage)
}

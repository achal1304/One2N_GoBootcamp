package main

const goalScore = 100

type Player struct {
	name            string
	minStrategy     int
	maxStrategy     int
	currentStrategy int
	dice            DiceRandom
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
		PrintOutcome(constantStrategy(fixedStrategy, variableStrategy))
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
		PrintMultipleGamesAtOnce(i, player1WinsSingleStrategy, (counter)*10)
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

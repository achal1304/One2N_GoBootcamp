package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func getStrategy(value string) int {
	parts := strings.Split(value, "-")
	if len(parts) == 1 {
		strategy, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}
		return strategy
	} else if len(parts) == 2 {
		return 0
	} else {
		fmt.Print("Incorrect format")
		return -1
	}
}

func main() {

	var player1Strategy string
	var player2Strategy string

	flag.StringVar(&player1Strategy, "player1", "", "Player 1 strategy")
	flag.StringVar(&player2Strategy, "player2", "", "Player 2 strategy")
	flag.Parse()

	strategy1 := getStrategy(player1Strategy)
	strategy2 := getStrategy(player2Strategy)

	if strategy1 < 0 || strategy2 < 0 {
		fmt.Println("Abandoned the game as strategy is incorrect")
	}
	if strategy1 > 0 && strategy2 > 0 {
		printOutcome(constantStrategy(strategy1, strategy2))
	} else if strategy1 > 0 && strategy2 == 0 {
		constantAndVariableStrategy(strategy1, strategy2)
	} else {
		variableAndVariableStrategy()
	}
}

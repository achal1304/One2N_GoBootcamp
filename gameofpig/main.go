package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func parseRange(value string) (int, int, error) {
	min, max := 0, 0

	strategyNumbers := strings.Split(value, "-")
	if len(strategyNumbers) != 2 {
		if len(strategyNumbers) == 1 {
			val, err := strconv.Atoi(strategyNumbers[0])
			if err != nil {
				return 0, 0, fmt.Errorf("invalid strategy format: %d", val)
			}
			return val, val, nil
		}
		return 0, 0, fmt.Errorf("invalid range format: %s", value)
	}

	min, err := strconv.Atoi(strategyNumbers[0])
	if err != nil || min < 0 {
		return 0, 0, fmt.Errorf("invalid min value: %s", strategyNumbers[0])
	}

	max, err = strconv.Atoi(strategyNumbers[1])
	if err != nil || max < 0 || max < min {
		return 0, 0, fmt.Errorf("invalid max value: %s", strategyNumbers[1])
	}

	return min, max, nil
}

func main() {

	var player1Strategy string
	var player2Strategy string

	flag.StringVar(&player1Strategy, "player1", "", "Player 1 strategy")
	flag.StringVar(&player2Strategy, "player2", "", "Player 2 strategy")
	flag.Parse()

	strategyMin1, strategyMax1, err := parseRange(player1Strategy)
	if err != nil {
		fmt.Println(err)
		return
	}
	strategyMin2, strategyMax2, err := parseRange(player2Strategy)
	if err != nil {
		fmt.Println(err)
		return
	}
	player1 := &Player{
		name:        "player1",
		minStrategy: strategyMin1,
		maxStrategy: strategyMax1,
		dice:        Dice{rng: rand.New(rand.NewSource(time.Now().UnixNano()))},
	}
	player2 := &Player{
		name:        "player2",
		minStrategy: strategyMin2,
		maxStrategy: strategyMax2,
		dice:        Dice{rng: rand.New(rand.NewSource(time.Now().UnixNano()))},
	}
	if strategyMin1 == strategyMax1 && strategyMin2 == strategyMax2 {
		player1.currentStrategy = player1.minStrategy
		player2.currentStrategy = player2.minStrategy
		PrintOutcome(constantStrategy(player1, player2))
	} else if strategyMin1 != strategyMax1 && strategyMin2 != strategyMax2 {
		variableAndVariableStrategy(player1, player2)
	} else {
		constantAndVariableStrategy(player1, player2)
	}
}

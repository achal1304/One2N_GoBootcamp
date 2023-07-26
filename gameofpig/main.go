package main

import (
	"flag"
)

func main() {

	var player1Strategy int
	var player2Strategy int

	flag.IntVar(&player1Strategy, "player1", 0, "Player 1 strategy")
	flag.IntVar(&player2Strategy, "player2", 0, "Player 2 strategy")
	flag.Parse()

	constantStrategy(player1Strategy, player2Strategy)
	// fmt.Println(player1Strategy, player2Strategy)
}

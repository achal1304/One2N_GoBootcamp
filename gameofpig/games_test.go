package main

import (
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/gameofpig/testutility"
)

func TestStrategyOutcome(t *testing.T) {
	testDice1 := testutility.NewTestDice([]int{1, 2, 3, 4, 5, 6})
	testDice2 := testutility.NewTestDice([]int{1, 2, 3, 4, 5, 6})

	testPlayer := &Player{
		currentStrategy: 10,
		dice:            testDice1,
	}
	testPlayer2 := &Player{
		dice:            testDice2,
		currentStrategy: 5,
	}

	testcases := []struct {
		testplayer     *Player
		expectedOutput int
	}{
		{testPlayer, 0},
		{testPlayer, 14},
		{testPlayer, 0},
		{testPlayer2, 0},
		{testPlayer2, 5},
		{testPlayer2, 9},
		{testPlayer2, 6},
	}

	for _, tc := range testcases {
		output := strategyOutcome(tc.testplayer)
		if output != tc.expectedOutput {
			t.Errorf("Expected %d got %d", tc.expectedOutput, output)
		}
	}
}

func TestConstantStrategy(t *testing.T) {
	testDice1 := testutility.NewTestDice([]int{1, 2, 3, 4, 5, 6})
	testDice2 := testutility.NewTestDice([]int{1, 2, 3, 4, 5, 6})

	testPlayer := &Player{
		currentStrategy: 10,
		dice:            testDice1,
	}
	testPlayer2 := &Player{
		dice:            testDice2,
		currentStrategy: 5,
	}
	testPlayer3 := &Player{
		dice:            testDice2,
		currentStrategy: 100,
	}
	testCases := []struct {
		player1          *Player
		player2          *Player
		player1Strat     int
		player2Strat     int
		player1TotalWins int
		totalGames       int
	}{
		{testPlayer, testPlayer2, 10, 5, 10, 10},
		{testPlayer2, testPlayer, 5, 10, 0, 10},
		{testPlayer2, testPlayer3, 5, 100, 10, 10},
	}

	for _, tc := range testCases {
		play1strat, play2strat, totalwins, totalgames := constantStrategy(tc.player1, tc.player2)
		if play1strat != tc.player1Strat {
			t.Errorf("Expected %d got %d", tc.player1Strat, play1strat)
		}
		if play2strat != tc.player2Strat {
			t.Errorf("Expected %d got %d", tc.player2Strat, play2strat)
		}
		if totalwins != tc.player1TotalWins {
			t.Errorf("Expected %d got %d", tc.player1TotalWins, totalwins)
		}
		if totalgames != tc.totalGames {
			t.Errorf("Expected %d got %d", tc.totalGames, totalgames)
		}
	}
}

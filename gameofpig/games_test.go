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

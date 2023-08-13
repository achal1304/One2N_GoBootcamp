package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRoll(t *testing.T) {
	testDice := Dice{rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
	for i := 0; i < 10; i++ {
		output := testDice.Roll()
		if output <= 0 || output > 6 {
			fmt.Errorf("Output should be between 1 and 6")
		}
	}
}

package main

import (
	"math/rand"
)

type DiceRandom interface {
	Roll() int
}

type Dice struct {
	rng *rand.Rand
}

func (d Dice) Roll() int {
	return d.rng.Intn(6) + 1
}

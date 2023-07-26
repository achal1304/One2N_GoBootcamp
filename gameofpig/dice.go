package main

import (
	"math/rand"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

type Dice struct {
}

func (d Dice) Roll() int {
	return rand.Intn(6) + 1
}

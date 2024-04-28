package bowling

import (
	"fmt"
)

const testVersion = 1

type Game struct {
	round     int
	prevPins  int
	data      []int
	maxRounds int
}

// NewGame sets up a new bowling game
func NewGame() *Game {
	return &Game{
		round:     0,
		prevPins:  -1,
		data:      make([]int, 0),
		maxRounds: 10,
	}
}

// Roll lets the game know how many pins where knocked down
func (g *Game) Roll(pins int) error {
	if g.gameOver() {
		return fmt.Errorf("Cannot roll after game is over %d", g.round)
	}

	if pins > 10 || pins < 0 {
		return fmt.Errorf("cannot knock down more than 10 pins or less than zero: %d", pins)
	}

	if g.prevPins != -1 && pins+g.prevPins > 10 {
		return fmt.Errorf("cannot knock down more than 10 pins in a round: %d + %d", g.prevPins, pins)
	}
	var spare bool
	if g.prevPins == -1 {
		g.prevPins = pins
		if pins == 10 {
			g.round++
			g.prevPins = -1
		} else if g.round >= 10 {
			g.round++
		}

	} else {
		spare = g.prevPins+pins == 10
		g.round++
		g.prevPins = -1
	}
	g.data = append(g.data, pins)

	if g.round == 10 {
		// last frame is a strike
		if pins == 10 && g.prevPins == -1 {
			g.maxRounds += 2
		} else if spare {
			g.maxRounds++
		}
	}

	return nil
}

func (g *Game) gameOver() bool {
	return g.round >= g.maxRounds
}

// Score the finished bowling game
func (g *Game) Score() (int, error) {
	if !g.gameOver() {
		return 0, fmt.Errorf("Game has not finished. Keep Bowling! %d %d", g.round, g.maxRounds)
	}
	score := 0
	prev := -1
	round := 0
	for i, h := range g.data {
		if prev == -1 {
			if h == 10 { // strike
				round++
				score += h + g.data[i+1] + g.data[i+2]
			} else {
				prev = h
			}
		} else {
			if h+prev == 10 { //spare
				score += g.data[i+1]
			}
			score += h + prev
			prev = -1
			round++
		}
		if round == 10 {
			break
		}

	}

	return score, nil
}

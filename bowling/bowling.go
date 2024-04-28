package bowling

import (
	"errors"
)

// Game represents a game of bowling.
type Game struct {
	frames        [12]int
	frameResult   [12]string
	frameIndex    int
	throwsInFrame int
	rolls         [12][2]int
}

// NewGame returns a pointer to a new game object.
func NewGame() *Game {
	return &Game{}
}

// Roll records a roll for a game of bowling.
func (g *Game) Roll(pins int) error {
	if pins < 0 {
		return errors.New("negative pins are not allowed")
	}
	if pins > 10 {
		return errors.New("pins greater than 10 are not allowed")
	}
	if 10-g.frames[g.frameIndex] < pins {
		return errors.New("cannot score more pins than are remaining")
	}
	if g.frameIndex == 10 && g.frames[9] != 10 {
		return errors.New("cannot roll after game is over")
	}
	if g.frameIndex == 10 && g.frameResult[9] == "-" && g.throwsInFrame == 1 {
		return errors.New("cannot roll after the game is over")
	}
	if g.frameResult[9] == "X" && g.frameResult[10] == "X" && g.frameResult[11] == "X" {
		return errors.New("cannot roll after game is over")
	}
	if g.frameResult[9] == "X" && g.frameResult[10] == "X" && g.throwsInFrame > 0 {
		return errors.New("cannot roll after game is over")
	}
	if g.frameResult[9] == "X" && g.frameResult[10] != "X" && g.frameIndex > 10 {
		return errors.New("cannot roll after game is over")
	}
	if pins == 10 {
		g.frameResult[g.frameIndex] = "X"
	} else if g.frames[g.frameIndex]+pins == 10 {
		g.frameResult[g.frameIndex] = "-"
	}
	g.rolls[g.frameIndex][g.throwsInFrame] = pins
	g.frames[g.frameIndex] += pins
	if g.throwsInFrame == 1 || pins == 10 {
		g.frameIndex++
		g.throwsInFrame = 0
	} else {
		g.throwsInFrame++
	}
	return nil
}

// Score returns the total score for a game of bowling.
func (g *Game) Score() (int, error) {
	if g.frameIndex < 9 {
		return 0, errors.New("score cannot be taken before the end of the game")
	}
	if g.frameResult[9] == "-" && g.frameIndex < 11 && g.throwsInFrame == 0 {
		return 0, errors.New("bonus roll for spare must be scored before end of game")
	}
	if g.frameResult[9] == "X" && g.frameIndex < 11 {
		return 0, errors.New("bonus rolls for strike must be scored before end of game")
	}
	if g.frameResult[9] == "X" && g.frameResult[10] == "X" && g.throwsInFrame == 0 && g.frameIndex < 12 {
		return 0, errors.New("bonus rolls for strike must be scored before end of game")
	}
	totalScore := 0
	for index, val := range g.frames {
		if g.frameResult[index] == "-" && index < 9 {
			totalScore += g.rolls[index+1][0]
		}
		if g.frameResult[index] == "X" && index < 9 {
			if g.frameResult[index+1] == "X" && g.frameResult[index+2] == "X" {
				totalScore += 20
			} else if g.frameResult[index+1] == "X" {
				totalScore += (10 + g.rolls[index+2][0])
			} else {
				totalScore += g.frames[index+1]
			}
		}
		totalScore += val
	}
	return totalScore, nil
}

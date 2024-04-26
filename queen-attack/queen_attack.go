package queenattack

import "errors"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == "" || blackPosition == "" || whitePosition == blackPosition {
		return false, errors.New("no correct input")
	}
	rWP := []rune(whitePosition)
	rBP := []rune(blackPosition)
	rWP0 := rWP[0] - 'a'
	rWP1 := rWP[1] - '1'
	rBP0 := rBP[0] - 'a'
	rBP1 := rBP[1] - '1'
	if 0 > rWP0 || rWP0 >= 8 || 0 > rWP1 || rWP1 >= 8 || 0 > rBP0 || rBP0 >= 8 || 0 > rBP1 || rBP1 >= 8 {
		return false, errors.New("no correct input")
	}
	if rWP0 == rBP0 || rWP1 == rBP1 || rWP0-rBP0 == rWP1-rBP1 || rWP0-rBP0 == rBP1-rWP1 || rBP0-rWP0 == rBP1-rWP1 {
		return true, nil
	}
	return false, nil
}

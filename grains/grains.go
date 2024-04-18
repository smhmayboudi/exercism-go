package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || 64 < number {
		return 0, errors.New("error")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return 0b1111111111111111111111111111111111111111111111111111111111111111
}

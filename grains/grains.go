package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number < 1 || 64 < number {
		return 0, errors.New("error")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}

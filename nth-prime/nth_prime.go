package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (res int, err error) {
	if n < 1 {
		return 0, errors.New("no less than 1")
	}
	if n == 1 {
		return 2, nil
	}
	if n == 2 {
		return 3, nil
	}
	idx := 2
	for i := idx; idx < n+2; i++ {
		if checkPrimeNumber(i) {
			res = i
			idx++
		}
	}
	return
}

func checkPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

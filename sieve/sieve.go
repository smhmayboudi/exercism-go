package sieve

func Sieve(limit int) []int {
	primes := make(map[int]bool, limit-2)
	for i := 0; i < limit-2; i++ {
		currentNumber := i + 2
		if primes[currentNumber] {
			continue
		}
		for j := 2; currentNumber*j <= limit; j++ {
			primes[currentNumber*j] = true
		}
	}
	res := []int{}
	for i, v := range primes {
		if !v {
			res = append(res, i)
		}
	}
	return res
}

package sieve

func Sieve(limit int) []int {
	cap := limit - 1
	strNum := 2
	primes := make(map[int]bool, cap)
	for i := 0; i < cap; i++ {
		if primes[i] {
			continue
		} else {
			primes[i] = false
		}
		currentNumber := i + strNum
		for j := currentNumber; j < cap; j = j + currentNumber {
			primes[j] = true
		}
	}
	res := []int{}
	for i, v := range primes {
		if !v {
			res = append(res, i+strNum)
		}
	}
	return res
}

package sieve

func Sieve(limit int) []int {
	composite := make([]bool, limit+1)
	primes := make([]int, limit/2)
	primeIndex := 0

	for number := 2; number <= limit; number++ {
		if !composite[number] {
			primes[primeIndex] = number
			primeIndex++
			for idx := number + number; idx <= limit; idx += number {
				composite[idx] = true
			}
		}
	}
	return primes[:primeIndex]
}

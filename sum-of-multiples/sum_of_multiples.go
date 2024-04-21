package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
Main:
	for i := 0; i < limit; i++ {
		for j := 0; j < len(divisors); j++ {
			divisor := divisors[j]
			if divisor != 0 && i%divisor == 0 {
				sum += i
				continue Main
			}
		}
	}
	return
}

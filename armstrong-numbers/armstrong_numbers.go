package armstrong

// IsNumber determines whether a number is an Armstrong number.
func IsNumber(number int) bool {
	exponent := 0
	for n := number; n > 0; n /= 10 {
		exponent++
	}
	sum := 0
	for n := number; n > 0; n /= 10 {
		digit := n % 10
		power := 1
		for j := 0; j < exponent; j++ {
			power *= digit
		}
		sum += power
	}
	return number == sum
}

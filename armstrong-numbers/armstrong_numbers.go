package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(number int) bool {
	numString := []rune(strconv.Itoa(number))
	numDigits := len(numString)
	sum := 0
	for _, value := range numString {
		num, _ := strconv.Atoi(string(value))
		sum += int(math.Pow(float64(num), float64(numDigits)))
	}
	return sum == number
}

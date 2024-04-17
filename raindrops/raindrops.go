package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	resule := ""
	if number%3 == 0 {
		resule += "Pling"
	}
	if number%5 == 0 {
		resule += "Plang"
	}
	if number%7 == 0 {
		resule += "Plong"
	}
	if !(number%3 == 0 || number%5 == 0 || number%7 == 0) {
		resule += strconv.Itoa(number)
	}
	return resule
}

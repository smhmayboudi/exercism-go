package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	resule := make([]string, 3)
	// switch number := number; {
	// case number%3 == 0:
	// 	resule = append(resule, "Pling")
	// case number%5 == 0:
	// 	resule = append(resule, "Pling")
	// case number%7 == 0:
	// 	resule = append(resule, "Plong")
	// }
	if number%3 == 0 {
		resule = append(resule, "Pling")
	}
	if number%5 == 0 {
		resule = append(resule, "Plang")
	}
	if number%7 == 0 {
		resule = append(resule, "Plong")
	}
	if !(number%3 == 0 || number%5 == 0 || number%7 == 0) {
		resule = append(resule, strconv.Itoa(number))
	}
	return strings.Join(resule, "")
}

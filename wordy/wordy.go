package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`-?\d+`)
	a := re.FindStringSubmatch(question)
	b := re.FindAllStringSubmatch(question, -1)
	plus, _ := regexp.MatchString(`(-?\d+) plus (-?\d+)`, question)
	minus, _ := regexp.MatchString(`-?\d+ minus -?\d+`, question)
	multipliedBy, _ := regexp.MatchString(`-?\d+ multiplied by -?\d+`, question)
	dividedBy, _ := regexp.MatchString(`-?\d+ divided by -?\d+`, question)

	fmt.Println(a, b, plus, minus, multipliedBy, dividedBy)

	total := 0
	for i := 0; i < len(b); i++ {
		v, _ := strconv.Atoi(b[i][0])
		if plus {
			total += v
		} else if minus {
			total -= v
		} else if multipliedBy {
			total *= v
		} else if dividedBy {
			total /= v
		} else {
			return 0, false
		}
	}
	return total, true
}

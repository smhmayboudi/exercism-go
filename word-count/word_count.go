package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frq := Frequency{}
	re := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
	res0 := re.FindAllString(strings.ToLower(phrase), -1)
	for i := 0; i < len(res0); i++ {
		frq[res0[i]] += 1
	}
	return frq
}

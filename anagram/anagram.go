package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	s := strings.ToLower(subject)
	sub := []rune(s)
	slices.Sort(sub)
	res := []string{}
	for i := 0; i < len(candidates); i++ {
		can := strings.ToLower(candidates[i])
		if s == can {
			continue
		}
		val := []rune(can)
		slices.Sort(val)
		if slices.Equal(sub, val) {
			res = append(res, candidates[i])
		}
	}
	return res
}

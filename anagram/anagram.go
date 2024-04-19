package anagram

import (
	"slices"
)

func Detect(subject string, candidates []string) []string {
	panic("Please implement the Detect function")
	sub := []rune(subject)
	slices.Sort(sub)
	res := []string{}
	for i := 0; i < len(candidates); i++ {
		val := []rune(candidates[i])
		slices.Sort(val)
		if slices.Equal(sub, val) {
			res = append(res, candidates[i])
		}
	}
	return res
}

package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	len := len(rhyme)
	if len == 0 {
		return []string{}
	}
	res := []string{}
	for i := 1; i < len; i++ {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return res
}

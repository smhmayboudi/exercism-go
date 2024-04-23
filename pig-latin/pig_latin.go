package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(text string) string {
	words := strings.Split(text, " ")
	res := make([]string, 0, len(words))
	for _, w := range words {
		res = append(res, word(w))
	}
	return strings.Join(res, " ")
}

var (
	beginVowelRe, _     = regexp.Compile(`^([aeiou].|xr|yt).+$`)
	beginConsonantRe, _ = regexp.Compile(`^([^aeiou]+)(.+)$`)
	quRe, _             = regexp.Compile(`^([^aeiou]*qu)(.+)$`)
	yVowelRe, _         = regexp.Compile(`^([^aeiou]+|.)(y.+)$`)
)

func word(word string) string {
	// Rule 1
	if beginVowelRe.MatchString(word) {
		return word + "ay"
	}
	// Rule 3
	if match := quRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	// Rule 4
	if match := yVowelRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	// Rule 2
	if match := beginConsonantRe.FindStringSubmatch(word); match != nil {
		return match[2] + match[1] + "ay"
	}
	return word
}

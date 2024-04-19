package atbash

import "strings"

func Atbash(s string) string {
	return strings.Map(func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return []rune("zyxwvutsrqponmlkjihgfedcba")[r-'a']
		} else if 'A' <= r && r <= 'Z' {
			return []rune("zyxwvutsrqponmlkjihgfedcba")[r-'A']
		}
		return r
	}, s)
}

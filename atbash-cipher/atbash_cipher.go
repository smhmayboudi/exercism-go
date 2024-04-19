package atbash

import "strings"

func Atbash(s string) string {
	var buf strings.Builder
	var cipher = []rune("zyxwvutsrqponmlkjihgfedcba")
	for _, c := range strings.ToLower(s) {
		if 'a' <= c && c <= 'z' {
			buf.WriteRune(cipher[c-'a'])
		} else if '0' <= c && c <= '9' {
			buf.WriteRune(c)
		}
		if buf.Len()%6 == 5 {
			buf.WriteRune(' ')
		}
	}
	return strings.TrimSpace(buf.String())
}

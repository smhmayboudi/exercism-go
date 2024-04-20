package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	str := []rune(s)
	var sb strings.Builder
	sb.Grow(len(str) / 2)
	for i := 0; i < len(str); i++ {
		if i == 0 || (1 < i && (str[i-1] == ' ' || str[i-1] == '-' || str[i-1] == '_')) {
			if unicode.IsLower(str[i]) {
				sb.WriteRune(str[i] - 'a' + 'A')
			} else if unicode.IsUpper(str[i]) {
				sb.WriteRune(str[i])
			}
		}
	}
	return sb.String()
}

package acronym

import "strings"

func Abbreviate(s string) string {
	str := []rune(s)
	var sb strings.Builder
	sb.Grow(len(str) / 2)
	for i := 0; i < len(str); i++ {
		if i == 0 || (1 < i && (str[i-1] == ' ' || str[i-1] == '-' || str[i-1] == '_')) {
			if 'a' <= str[i] && str[i] <= 'z' {
				sb.WriteRune(str[i] - 'a' + 'A')
			} else if 'A' <= str[i] && str[i] <= 'Z' {
				sb.WriteRune(str[i])
			}
		}
	}
	return sb.String()
}

package rotationalcipher

import "strings"

func RotationalCipher(input string, shift int) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(shift))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(shift))%26
		}
		return r
	}, input)
}

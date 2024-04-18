package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	result := uint32(0)
	for _, value := range strings.ToLower(input) {
		switch value := value; {
		case value == ' ', value == '_', value == '.', value == '"', '0' <= value && value <= '9':
			continue
		case value < 'a' || 'z' < value:
			return false
		default:
			result |= 1 << (value - 'a')
		}
	}
	return result == 0x3ffffff
}

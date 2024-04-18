package pangram

func IsPangram(input string) bool {
	result := uint32(0)
	for _, value := range input {
		switch value := value; {
		case value == ' ', value == '_', value == '.', value == '"', '0' <= value && value <= '9':
			continue
		case 'a' <= value && value <= 'z':
			result |= 1 << (value - 'a')
		case 'A' <= value && value <= 'Z':
			result |= 1 << (value - 'A')
		}
	}
	return result == 0b11111111111111111111111111
}

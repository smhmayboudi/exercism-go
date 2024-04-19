package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey < 1 || 25 < shiftKey {
		return plain
	}
	newPlain := []rune(plain)
	for i := 0; i < len(newPlain); i++ {
		if 'a' <= newPlain[i] && newPlain[i] <= 'z' {
			newPlain[i] = 'a' + (newPlain[i]-'a'+rune(shiftKey))%26
		}
		if 'A' <= newPlain[i] && newPlain[i] <= 'Z' {
			newPlain[i] = 'A' + (newPlain[i]-'A'+rune(shiftKey))%26
		}
	}
	return string(newPlain)
}

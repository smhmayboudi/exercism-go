package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey < 1 || 25 < shiftKey {
		return plain
	}
	newPlain := []rune(plain)
	for i := 0; i < len(newPlain); i++ {
		check0 := newPlain[i]
		if check0 == ' ' {
			continue
		}
		check1 := check0 + rune(shiftKey)
		if 'a' <= check0 && check0 <= 'z' {
			if check1 > 'z' {
				newPlain[i] = check1 - rune(26)
			} else {
				newPlain[i] = check1
			}
		}
		if 'A' <= check0 && check0 <= 'Z' {
			if check1 > 'Z' {
				newPlain[i] = check1 - rune(26)
			} else {
				newPlain[i] = check1
			}
		}
	}
	return string(newPlain)
}

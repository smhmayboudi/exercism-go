package rotationalcipher

import "fmt"

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
		a, b, c, d := 'a', 'z', 'A', 'Z'
		fmt.Println(a, b, c, d)
		if ('a' <= check0 && check0 <= 'z') ||
			('A' <= check0 && check0 <= 'Z') {
			if ('a' <= check0 && check0 <= 'z' && 'a' <= check1 && check1 <= 'z') ||
				('A' <= check0 && check0 <= 'Z' && 'A' <= check1 && check1 <= 'Z') {
				newPlain[i] = check1
			} else {
				newPlain[i] = newPlain[i] - rune(shiftKey)
			}
		}
	}
	return string(newPlain)
}

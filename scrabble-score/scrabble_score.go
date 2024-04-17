package scrabble

import "unicode"

func Score(word string) (sum int) {
	for _, letter := range word {
		switch unicode.ToLower(letter) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			sum += 1
		case 'd', 'g':
			sum += 2
		case 'b', 'c', 'm', 'p':
			sum += 3
		case 'f', 'h', 'v', 'w', 'y':
			sum += 4
		case 'k':
			sum += 5
		case 'j', 'x':
			sum += 8
		case 'q', 'z':
			sum += 10
		}
	}
	return sum
}

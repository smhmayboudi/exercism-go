package piglatin

import "strings"

// isVowel returns true for a,e,i,o,u.
func isVowel(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }

// findVowel finds the first vowel sound in the word.
func findVowel(word string) int {
	for i, c := range []byte(word) {
		if isVowel(c) && (i == 0 || word[i-1] != 'q' || c != 'u') {
			return i
		} else if (c == 'y' || c == 'x') && (i == len(word)-1 || !isVowel(word[i+1])) {
			return i
		}
	}
	return len(word)
}

// Sentence converts message to pig latin
func Sentence(input string) string {

	// For each word in input
	words := strings.Fields(input)
	for i, word := range words {

		// Find the first j sound, and modify word
		j := findVowel(word)
		words[i] = word[j:] + word[:j] + "ay"
	}
	return strings.Join(words, " ")
}

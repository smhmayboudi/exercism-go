package isogram

import "strings"

func IsIsogram(word string) bool {
	hashset := map[rune]struct{}{}
	newWord := strings.ToLower(word)
	newWord = strings.Replace(newWord, " ", "", -1)
	newWord = strings.Replace(newWord, "-", "", -1)
	for _, value := range newWord {
		_, b := hashset[value]
		if b {
			return false
		}
		hashset[value] = struct{}{}
	}
	return true
}

package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnRune := []rune(isbn)
	ln := len(isbnRune)
	if ln != 10 {
		return false
	}
	sum := 0
	for idx := 0; idx < ln; idx++ {
		val := isbnRune[idx]
		if val == 'X' {
			if idx == ln-1 {
				valI := 10
				sum += valI * (10 - idx)
			} else {
				return false
			}
		} else if '0' <= val && val <= '9' {
			valI, _ := strconv.Atoi(string(val))
			sum += valI * (10 - idx)
		} else {
			return false
		}
	}
	return sum%11 == 0
}

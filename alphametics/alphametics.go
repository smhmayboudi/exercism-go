package alphametics

import (
	"errors"
	"strings"
)

type character struct {
	val     int
	letter  byte
	leading bool // Letter appears at start of word (cannot be '0')
}

func Solve(puzzle string) (map[string]int, error) {

	var usedDigits [10]bool
	var letters ['Z']character // Array of each possible letter, to avoid using a map

	// Split string into two sides, then operator and operands
	operands := strings.FieldsFunc(puzzle, func(r rune) bool { return r == '+' || r == '=' || r == ' ' })

	// Flag used letters, and which appear at the start of words
	for _, w := range operands {
		letters[w[0]].leading = true
		for _, c := range []byte(w) {
			letters[c].letter = c
		}
	}

	// Create an index into the array so it's easier to allocate to the nth letter later
	index := make([]*character, 0, 26)
	for i := range letters['A':'Z'] {
		if letters[i+'A'].letter > 0 {
			index = append(index, &letters[i+'A'])
		}
	}

	// Recursive helper for Solve
	var solveInner func(int) bool
	solveInner = func(depth int) bool {

		// If we've allocated every letter, calculate if it satisfies sum
		if depth == len(index) {
			tot := 0
			for _, operand := range operands[:len(operands)-1] {
				tot += wordVal(letters[:], operand)
			}
			return tot == wordVal(letters[:], operands[len(operands)-1])
		}

		// Otherwise, try every value for the next letter
		for char, digit := index[depth], 0; digit < 10; digit++ {
			if (digit != 0 || !char.leading) && !usedDigits[digit] {
				usedDigits[digit], char.val = true, digit
				if solveInner(depth + 1) {
					return true
				}
				usedDigits[digit] = false
			}
		}
		return false
	}

	// Use recursion to try all values for each letter
	if !solveInner(0) {
		return nil, errors.New("no solution")
	}

	// Format successful result
	result := make(map[string]int, len(index))
	for _, v := range index {
		result[string(v.letter)] = v.val
	}
	return result, nil
}

// wordVal returns the value of the supplied word for these letter values
func wordVal(letters []character, word string) (sum int) {
	for _, c := range []byte(word) {
		sum = 10*sum + letters[c].val
	}
	return sum
}

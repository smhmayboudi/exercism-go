package hamming

import "errors"

func Distance(a, b string) (int, error) {
	a2 := []rune(a)
	b2 := []rune(b)
	if len(a2) != len(b2) {
		return 0, errors.New("inputs should have same length")
	}
	count := 0
	for i := 0; i < len(a2); i++ {
		if a2[i] != b2[i] {
			count++
		}
	}
	return count, nil
}

package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n < 1 {
		return 0, errors.New("input number should be more than 0")
	}
	count++
	if n == 1 {
		return 0, nil
	} else if n%2 == 0 {
		a, b := CollatzConjecture(n / 2)
		return a + count, b
	} else if n%2 == 1 {
		a, b := CollatzConjecture(3*n + 1)
		return a + count, b
	}
	return count, nil
}

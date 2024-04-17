package collatzconjecture

import "errors"

func CollatzConjecture(n int) (i int, err error) {
	if n < 1 {
		return 0, errors.New("input number should be more than 0")
	}
	for i := 0; n != 1; i++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return
}

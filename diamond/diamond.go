package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' && 'Z' < char {
		return "", errors.New("invalid character.")
	}
	dim := int(char - 'A')
	if dim < 1 {
		return "A", nil
	}
	middle := dim
	dim = 2 * dim
	var sb strings.Builder
	sb.Grow(dim * dim)
	ch := byte('A')
	for r := 0; r <= dim; r++ {
		for c := 0; c <= dim; c++ {
			if r <= middle {
				if c == middle-r || c == middle+r {
					sb.WriteByte(ch)
				} else {
					sb.WriteByte(' ')
				}
			} else {
				if c == r-middle || c == 3*middle-r {
					sb.WriteByte(ch)
				} else {
					sb.WriteByte(' ')
				}
			}
		}
		if r < middle {
			ch += 1
		} else {
			ch -= 1
		}
		if r < dim {
			sb.WriteByte('\n')
		}
	}
	return sb.String(), nil
}

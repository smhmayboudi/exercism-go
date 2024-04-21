package cryptosquare

import (
	"math"
	"strings"
)

func Encode(pt string) string {
	runes := []rune(pt)
	var sb strings.Builder

	le := len(runes)
	sb.Grow(le)

	// Filter
	for i := 0; i < le; i++ {
		v := runes[i]
		if 'A' <= v && v <= 'Z' {
			sb.WriteRune(v - 'A' + 'a')
		} else if 'a' <= v && v <= 'z' {
			sb.WriteRune(v)
		} else if '0' <= v && v <= '9' {
			sb.WriteRune(v)
		}
	}

	// Rectangle
	pt = sb.String()
	runes = []rune(pt)
	le = len(runes)

	r, c := Box(le)

	sb.Reset()
	sb.Grow(le + r + c)

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			if i+j*c < le {
				sb.WriteRune(runes[i+j*c])
			} else {
				sb.WriteRune(' ')
			}
		}
		if i+1 < c {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func Box(le int) (r int, c int) {
	start := int(math.Sqrt(float64(le)))
	for row := start; row <= le/2+1; row++ {
		for col := start; col <= le/2+1; col++ {
			if le <= row*col && row <= col && col-row <= 1 {
				r = row
				c = col
				return
			}
		}
	}
	return
}

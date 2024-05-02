package palindrome

import (
	"errors"
	"math"
)

// Product represents a palidrome and all its possible two-factor factorizations.
type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalindrome(p int) bool {
	revp := 0
	for n := p; n > 0; n /= 10 {
		revp = revp*10 + n%10
	}
	return p == revp
}

// Products finds the largest and smallest palindromes which are products of numbers within given range.
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}
	pmin.Product = math.MaxInt64
	pmax.Product = math.MinInt64
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if (p <= pmin.Product || p >= pmax.Product) && isPalindrome(p) {
				switch {
				case p < pmin.Product:
					pmin = Product{p, [][2]int{{i, j}}}
				case p == pmin.Product:
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}
				switch {
				case p > pmax.Product:
					pmax = Product{p, [][2]int{{i, j}}}
				case p == pmax.Product:
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	if pmin.Product == math.MaxInt64 {
		return pmin, pmax, errors.New("no palindromes")
	}
	return pmin, pmax, nil
}

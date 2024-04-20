package listops

import "slices"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := 0; i < len(s); i++ {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; 0 <= i; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := []int{}
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			out = append(out, s[i])
		}
	}
	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i := 0; i < len(s); i++ {
		s[i] = fn(s[i])
	}
	return s
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	a := slices.Concat(lists...)
	return s.Append(a)
}

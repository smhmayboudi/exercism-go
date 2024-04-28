package twobucket

import (
	"errors"
	"math"
)

type state [2]int

func Solve(size1 int, size2 int, goal int, start string) (string, int, int, error) {
	if size1 == 0 || size2 == 0 || goal == 0 || (start != "one" && start != "two") {
		return "", 0, 0, errors.New("Invalid data")
	}

	handled := make(map[state]bool)
	var last []state
	if start == "one" {
		last = []state{state{size1, 0}}
	} else {
		last = []state{state{0, size2}}
	}

	cnt := 1
	for {
		for _, n := range last {
			if n[0] == goal {
				return "one", cnt, n[1], nil
			}
			if n[1] == goal {
				return "two", cnt, n[0], nil
			}
		}

		current := []state{}
		for _, state := range last {
			perm := permuts(size1, state[0], size2, state[1])

			for _, p := range perm {
				if !handled[p] && isValid(p, size1, size2, start) {
					current = append(current, p)
				}
				handled[p] = true
			}
		}
		if len(current) == 0 {
			return "", 0, 0, errors.New("...")
		}
		cnt++
		last = current
	}
}

func isValid(s state, size1 int, size2 int, start string) bool {
	if s[0] < 0 || s[1] < 0 {
		return false
	}
	if start == "one" {
		return !(s[0] == 0 && s[1] == size2)
	} else {
		return !(s[0] == size1 && s[1] == 0)
	}
}

func permuts(size1 int, fill1 int, size2 int, fill2 int) []state {
	diff1 := int(math.Min(float64(size2-fill2), float64(fill1)))
	diff2 := int(math.Min(float64(size1-fill1), float64(fill2)))

	next := []state{
		state{fill1 - diff1, fill2 + diff1},
		state{fill1 + diff2, fill2 - diff2},
		state{0, fill2},
		state{fill1, 0},
		state{size1, fill2},
		state{fill1, size2},
	}
	return next
}
